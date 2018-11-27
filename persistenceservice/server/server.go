package server

import (
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
	return gorm.Open("postgres", connStr)
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
			ValueType:   "Binary",
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

	logging.LogInfo("Getting cache value: " + stringutil.PartialMessage(logRecord.Message))

	db, err := p.Configuration.GetConnection()
	if err != nil {
		return err
	}

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

func (p *PersistenceServiceEndpoint) GetBuildInfrastructure(key inframodel.RepositoryKey) (
	*inframodel.BuildInfrastructureMetadata, error) {
	logging.LogInfo("Getting infrastructure metadata")

	var im inframodel.BuildInfrastructureMetadata
	db, err := p.Configuration.GetConnection()
	if err != nil {
		return nil, err
	}

	resource := getResource(key, db)
	sites := getIisSites(resource.IisSites, db)
	allSites := getAllIisSites(db)
	applications := getIisApplications(resource.IisApplications, db)
	appPools := getRelevantAppPools(*sites, *applications)
	appPoolNames := getAppPoolNames(appPools)
	tasks := getWindowsTasks(resource.ScheduledTasks, db)
	taskNames := getScheduledTaskNames(*tasks)
	services := getWindowsServices(resource.WindowsServices, db)
	serviceNames := getWindowsServiceNames(*services)
	environments := getEnvironments(db)
	deploymentLocations := getDeploymentLocations(*applications, *sites, *tasks, *services)

	var results []inframodel.ServerTypeMetadata
	for _, environment := range *environments {
		// TODO: Move business logic somewhere else, seems un-ideal here
		siteInfo := getIisSiteForEnvironment(&environment, allSites)

		deploymentLocationsTrans := templating.TranscludeVariableInList(
			deploymentLocations,
			"SiteName",
			siteInfo.SiteName)

		appPoolNamesTrans := templating.TranscludeVariableInList(
			appPoolNames,
			"SiteName",
			siteInfo.SiteName)

		for _, server := range environment.Servers {
			if server.ServerType != resource.Type {
				continue
			}

			results = append(results, inframodel.ServerTypeMetadata{
				ServerName:          server.ServerName,
				ServerType:          server.ServerType,
				EnvironmentName:     environment.GetEnvironmentName(),
				DeploymentLocations: deploymentLocationsTrans,
				AppPoolNames:        appPoolNamesTrans,
				ServiceNames:        serviceNames,
				TaskNames:           taskNames,
			})
		}
	}

	im.Metadata = results
	return &im, nil
}
