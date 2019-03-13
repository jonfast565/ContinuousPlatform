// Persistence Service server
package server

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/databasemodels"
	"github.com/jonfast565/continuous-platform/models/inframodel"
	"github.com/jonfast565/continuous-platform/models/loggingmodel"
	"github.com/jonfast565/continuous-platform/models/persistmodel"
	"github.com/jonfast565/continuous-platform/utilities/cache"
	"github.com/jonfast565/continuous-platform/utilities/logging"
	"github.com/jonfast565/continuous-platform/utilities/networking"
	"github.com/jonfast565/continuous-platform/utilities/stringutil"
	"github.com/jonfast565/continuous-platform/utilities/templating"
	"github.com/satori/go.uuid"
	"os/user"
	"time"
)

// Persistence Service Configuration object
type PersistenceServiceConfiguration struct {
	Port      int
	Username  string
	Password  string
	Host      string
	DbPort    int
	DbName    string
	EnableSsl bool
}

// Creates a valid postgres connection string from a configuration object
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

// Gets a valid Gorm (ORM) connection from a Postgres connection string
func (c *PersistenceServiceConfiguration) GetConnection() (*gorm.DB, error) {
	connStr := c.GetPostgresConnectionString()
	return gorm.Open(constants.DatabaseDriver, connStr)
}

// Service endpoint
type PersistenceServiceEndpoint struct {
	Configuration PersistenceServiceConfiguration
}

// Service endpoint constructor
func NewPersistenceServiceEndpoint(configuration PersistenceServiceConfiguration) *PersistenceServiceEndpoint {
	result := new(PersistenceServiceEndpoint)
	result.Configuration = configuration
	return result
}

// Set the key value cache based on a request
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

// Get the key value cache based on a request
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

// Sets a log record in the database (for use with logging libraries and not directly)
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

// Gets a list of resources that can be extracted from the database (i.e. servers, apps, sites, etc.)
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

// Gets the infrastructure for a given resource
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
	// TODO: Remove, is it required
	// allSiteParts := getAllIisSiteParts(db)
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
	if len(*applications) > 0 {
		for _, environment := range *environments {
			for _, application := range *applications {
				siteInfos := getIisSitesForEnvironment(&environment, allSites, &application)
				var deploymentLocationsWithSiteName []string
				var appPoolsWithSiteName []string
				for _, siteInfo := range *siteInfos {
					deploymentLocationsWithSiteName = templating.TranscludeVariableInList(
						deploymentLocations,
						"SiteName",
						siteInfo.SiteName)

					appPoolsWithSiteName = templating.TranscludeVariableInList(
						appPoolNames,
						"SiteName",
						siteInfo.SiteName)

					for _, server := range environment.Servers {
						results = append(results, inframodel.ServerTypeMetadata{
							ServerName:          server.ServerName,
							EnvironmentName:     environment.GetEnvironmentName(),
							DeploymentLocations: deploymentLocationsWithSiteName,
							AppPoolNames:        appPoolsWithSiteName,
							ServiceNames:        serviceNames,
							TaskNames:           taskNames,
						})
					}
				}
			}
		}
		im.Metadata = results
		cache.SetCache("InfrastructureMetadata", key.String(), &im)
		return &im, nil
	}

	for _, environment := range *environments {
		for _, server := range environment.Servers {
			results = append(results, inframodel.ServerTypeMetadata{
				ServerName:          server.ServerName,
				EnvironmentName:     environment.GetEnvironmentName(),
				DeploymentLocations: deploymentLocations,
				AppPoolNames:        appPoolNames,
				ServiceNames:        serviceNames,
				TaskNames:           taskNames,
			})
		}
	}

	im.Metadata = results
	cache.SetCache("InfrastructureMetadata", key.String(), &im)
	return &im, nil
}
