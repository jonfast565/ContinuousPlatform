package main

import (
	"../databasemodels"
	"../jsonutil"
	"../logging"
	"./importmodels"
	"github.com/ahmetb/go-linq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satori/go.uuid"
)

func migrateSchema(db *gorm.DB) {
	logging.LogInfo("Migrating schema")
	db.AutoMigrate(
		&databasemodels.Environment{},
		&databasemodels.AppCache{},
		&databasemodels.Log{},
		&databasemodels.IisSite{},
		&databasemodels.IisApplication{},
		&databasemodels.IisApplicationPool{},
		&databasemodels.ScheduledTask{},
		&databasemodels.WindowsService{},
		&databasemodels.Resource{},
		&databasemodels.Server{})
}

func flushTables(db *gorm.DB) {
	logging.LogInfo("Flushing data")
	db.DropTable(&databasemodels.Environment{})
	db.DropTable(&databasemodels.AppCache{})
	db.DropTable(&databasemodels.IisSite{})
	db.DropTable(&databasemodels.IisApplication{})
	db.DropTable(&databasemodels.IisApplicationPool{})
	db.DropTable(&databasemodels.ScheduledTask{})
	db.DropTable(&databasemodels.WindowsService{})
	db.DropTable(&databasemodels.Resource{})
	db.DropTable(&databasemodels.Server{})
}

func loadData(i *importmodels.InfraImport, db *gorm.DB) {
	logging.LogInfo("Loading data")

	for _, iisApplicationPools := range i.IisApplicationPools {
		id, _ := uuid.NewV4()
		if result := db.Create(&databasemodels.IisApplicationPool{
			IisApplicationPoolId: id,
			Name:                 iisApplicationPools.Name,
			ProcessType:          iisApplicationPools.ProcessType,
			FrameworkVersion:     iisApplicationPools.FrameworkVersion,
		}); result.Error != nil {
			panic(result.Error)
		}
	}

	for _, iisApp := range i.IisApplications {
		id, _ := uuid.NewV4()
		var appPool databasemodels.IisApplicationPool
		if result := db.First(&appPool,
			&databasemodels.IisApplicationPool{Name: iisApp.ApplicationPool}); result.Error != nil {
			panic(result.Error)
		}
		if result := db.Create(&databasemodels.IisApplication{
			IisApplicationId:  id,
			ApplicationName:   iisApp.Name,
			ApplicationPoolId: appPool.IisApplicationPoolId,
			PhysicalPath:      iisApp.PhysicalPath,
		}); result.Error != nil {
			panic(result.Error)
		}
	}

	for _, iisSite := range i.IisSites {
		id, _ := uuid.NewV4()

		var appPool databasemodels.IisApplicationPool
		if result := db.First(&appPool,
			&databasemodels.IisApplicationPool{Name: iisSite.ApplicationPool}); result.Error != nil {
			panic(result.Error)
		}

		var applications []databasemodels.IisApplication
		if result := db.Model(&databasemodels.IisApplication{}).
			Where("application_name in (?)", iisSite.Applications).
			Find(&applications); result.Error != nil {
			panic(result.Error)
		}

		var ids []string
		linq.From(applications).SelectT(func(application databasemodels.IisApplication) string {
			return application.IisApplicationId.String()
		}).ToSlice(&ids)

		if result := db.Create(&databasemodels.IisSite{
			IisSiteId:         id,
			SiteName:          iisSite.Name,
			ApplicationPoolId: appPool.IisApplicationPoolId,
			PhysicalPath:      iisSite.PhysicalPath,
			SiteApplications:  ids,
		}); result.Error != nil {
			panic(result.Error)
		}
	}

	for _, windowsService := range i.WindowsServices {
		id, _ := uuid.NewV4()
		if result := db.Create(&databasemodels.WindowsService{
			WindowsServiceId:          id,
			Name:                      windowsService.Name,
			BinaryPath:                windowsService.BinaryPath,
			BinaryExecutableName:      windowsService.BinaryExecutableName,
			BinaryExecutableArguments: windowsService.BinaryExecutableArguments,
		}); result.Error != nil {
			panic(result.Error)
		}
	}

	for _, scheduledTask := range i.ScheduledTasks {
		id, _ := uuid.NewV4()
		if result := db.Create(&databasemodels.ScheduledTask{
			ScheduledTaskId:           id,
			Name:                      scheduledTask.Name,
			BinaryPath:                scheduledTask.BinaryPath,
			BinaryExecutableName:      scheduledTask.BinaryExecutableName,
			BinaryExecutableArguments: scheduledTask.BinaryExecutableArguments,
			ScheduleType:              scheduledTask.ScheduleType,
			RepeatInterval:            scheduledTask.RepeatInterval,
			RepetitionDuration:        scheduledTask.RepetitionDuration,
			ExecutionTimeLimit:        scheduledTask.ExecutionTimeLimit,
			Priority:                  scheduledTask.Priority,
		}); result.Error != nil {
			panic(result.Error)
		}
	}

	for _, server := range i.Servers {
		id, _ := uuid.NewV4()
		if result := db.Create(&databasemodels.Server{
			ServerId: id,
			Name:     server.Name,
			Type:     server.Type,
		}); result.Error != nil {
			panic(result.Error)
		}
	}

	for _, environment := range i.Environments {
		id, _ := uuid.NewV4()

		var servers []databasemodels.Server
		if result := db.Model(&databasemodels.Server{}).
			Where("name in (?)", environment.Servers).
			Find(&servers); result.Error != nil {
			panic(result.Error)
		}

		var ids []string
		linq.From(servers).SelectT(func(server databasemodels.Server) string {
			return server.ServerId.String()
		}).ToSlice(&ids)

		if result := db.Create(&databasemodels.Environment{
			EnvironmentId: id,
			Name:          environment.Name,
			BusinessLine:  environment.BusinessLine,
			Servers:       ids,
		}); result.Error != nil {
			panic(result.Error)
		}
	}
}

func main() {
	logging.LogHeader("Data Loader")
	logging.CreateLog()

	var c Configuration
	jsonutil.DecodeJsonFromFile("./appsettings.json", &c)

	connStr := c.GetPostgresConnectionString()
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// db.LogMode(true)

	flushTables(db)
	migrateSchema(db)

	var i importmodels.InfraImport
	jsonutil.DecodeJsonFromFile("./data.json", &i)
	loadData(&i, db)

	logging.LogInfo("Done")
}
