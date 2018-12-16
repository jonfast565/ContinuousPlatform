package server

import (
	"../../cache"
	"../../constants"
	"../../databasemodels"
	"../../logging"
	"../../models/inframodel"
	"../../models/loggingmodel"
	"../../models/persistmodel"
	"../../networking"
	"../../stringutil"
	"../../templating"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"os/user"
	"time"
)

type PersistenceServiceConfiguration struct {
	Port      int
	Username  string
	Password  string
	Host      string
	DbPort    int
	DbName    string
	EnableSsl bool
}

func (c PersistenceServiceConfiguration) GetPostgresConnectionString() string {
	enableSslResult := "enable"
	if c.EnableSsl == false {
		enableSslResult = "disable"
	}
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		c.Host,
		c.DbPort,
		c.Username,
		c.DbName,
		c.Password,
		enableSslResult)
}

func (c *PersistenceServiceConfiguration) GetConnection() (*gorm.DB, error) {
	connStr := c.GetPostgresConnectionString()
	return gorm.Open(constants.DatabaseDriver, connStr)
}

type PersistenceServiceEndpoint struct {
	Configuration PersistenceServiceConfiguration
}

func NewPersistenceServiceEndpoint(configuration PersistenceServiceConfiguration) *PersistenceServiceEndpoint {
	result := new(PersistenceServiceEndpoint)
	result.Configuration = configuration
	return result
}

func (p *PersistenceServiceEndpoint) SetKeyValueCache(
	setRequest *persistmodel.KeyValueRequest) error {

	logging.LogInfo("Setting cache value: " + setRequest.Key)

	db, err := p.Configuration.GetConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	hostname, err := networking.GetMyHostName()
	if err != nil {
		return err
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	var result databasemodels.AppCache
	cacheState := db.First(&result, &databasemodels.AppCache{
		KeyString:   setRequest.Key,
		MachineName: hostname,
	})

	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	if !cacheState.RecordNotFound() {
		db.Model(&result).Update("value", setRequest.Value)
		db.Model(&result).Update("last_modified_date_time", time.Now())
		db.Model(&result).Update("last_modified_by", currentUser.Name)
	} else {
		db.Create(&databasemodels.AppCache{
			AppCacheId:  uid,
			MachineName: hostname,
			KeyString:   setRequest.Key,
			Value:       setRequest.Value,
			ValueType:   constants.ValueType,
			AuditFields: databasemodels.AuditFields{
				CreatedDateTime:      time.Now(),
				CreatedBy:            currentUser.Name,
				LastModifiedDateTime: time.Now(),
				LastModifiedBy:       currentUser.Name,
			},
		})
	}

	return nil
}

func (p *PersistenceServiceEndpoint) GetKeyValueCache(
	getRequest *persistmodel.KeyRequest) (*persistmodel.KeyValueResult, error) {
	logging.LogInfo("Getting cache value: " + getRequest.Key)

	db, err := p.Configuration.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	hostname, err := networking.GetMyHostName()
	if err != nil {
		return nil, err
	}

	var result databasemodels.AppCache
	if db.First(&result, &databasemodels.AppCache{
		KeyString:   getRequest.Key,
		MachineName: hostname,
	}).RecordNotFound() {
		return nil, errors.New("value not found")
	}

	return &persistmodel.KeyValueResult{Value: result.Value}, nil
}

func (p *PersistenceServiceEndpoint) SetLogRecord(logRecord *loggingmodel.LogRecord) error {

	fmt.Printf("Getting cache value: %s", stringutil.PartialMessage(logRecord.Message))

	db, err := p.Configuration.GetConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	hostname, err := networking.GetMyHostName()
	if err != nil {
		return err
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	db.Create(&databasemodels.Log{
		LogId:           uid,
		MachineName:     hostname,
		ApplicationName: logRecord.ApplicationName,
		LogLevel:        logRecord.LogLevel,
		Message:         logRecord.Message,
		AuditFields: databasemodels.AuditFields{
			CreatedDateTime:      time.Now(),
			CreatedBy:            currentUser.Name,
			LastModifiedDateTime: time.Now(),
			LastModifiedBy:       currentUser.Name,
		},
	})
	return nil
}

func (p *PersistenceServiceEndpoint) GetResourceCache() (*inframodel.ResourceList, error) {
	logging.LogInfo("Getting resource cache")

	var result inframodel.ResourceList
	result.Keys = make([]inframodel.ResourceKey, 0)

	db, err := p.Configuration.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var resources []databasemodels.Resource
	if db.Find(&resources, databasemodels.Resource{}).RecordNotFound() {
		panic("resources not found")
	}

	for _, resource := range resources {
		result.Keys = append(result.Keys, inframodel.ResourceKey{
			RepositoryName: resource.RepositoryName,
			SolutionName:   resource.SolutionName,
			ProjectName:    resource.ProjectName,
		})
	}

	return &result, nil
}

func (p *PersistenceServiceEndpoint) GetBuildInfrastructure(key inframodel.ResourceKey) (
	*inframodel.BuildInfrastructureMetadata, error) {
	logging.LogInfo("Getting infrastructure metadata: " + key.String())

	if data, found := cache.GetCache("InfrastructureMetadata", key.String()); found == true {
		im := data.(*inframodel.BuildInfrastructureMetadata)
		return im, nil
	}

	var im inframodel.BuildInfrastructureMetadata
	db, err := p.Configuration.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	resource := getResource(key, db)
	sites := getIisSites(resource.IisSites, db)
	allSiteParts := getAllIisSiteParts(db)
	allSites := getAllIisSites(db)
	applications := getIisApplications(resource.IisApplications, db)
	appPools := getRelevantAppPools(*sites, *applications)
	appPoolNames := getAppPoolNames(appPools)
	tasks := getWindowsTasks(resource.ScheduledTasks, db)
	taskNames := getScheduledTaskNames(*tasks)
	services := getWindowsServices(resource.WindowsServices, db)
	serviceNames := getWindowsServiceNames(*services)
	environments := getEnvironments(db)
	environments = filterEnvironments(*applications, *allSites, *sites, *tasks, *services, *environments)
	deploymentLocations := getDeploymentLocations(*applications, *sites, *tasks, *services)

	var results []inframodel.ServerTypeMetadata
	for _, environment := range *environments {
		// TODO: Is this algorithm even correct?
		siteInfos := getIisSitesForEnvironment(&environment, allSiteParts)
		for _, siteInfo := range *siteInfos {
			var deploymentLocationsTrans []string
			var appPoolNamesTrans []string
			if len(*siteInfos) > 0 {
				deploymentLocationsTrans = templating.TranscludeVariableInList(
					deploymentLocations,
					"SiteName",
					siteInfo.SiteName)

				appPoolNamesTrans = templating.TranscludeVariableInList(
					appPoolNames,
					"SiteName",
					siteInfo.SiteName)
			} else {
				deploymentLocationsTrans = deploymentLocations
				appPoolNamesTrans = appPoolNames
			}

			for _, server := range environment.Servers {
				results = append(results, inframodel.ServerTypeMetadata{
					ServerName:          server.ServerName,
					EnvironmentName:     environment.GetEnvironmentName(),
					DeploymentLocations: deploymentLocationsTrans,
					AppPoolNames:        appPoolNamesTrans,
					ServiceNames:        serviceNames,
					TaskNames:           taskNames,
				})
			}
		}
	}

	im.Metadata = results
	cache.SetCache("InfrastructureMetadata", key.String(), &im)
	return &im, nil
}
