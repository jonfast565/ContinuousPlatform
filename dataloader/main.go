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

	for _, iisApp := range i.IisApplications {
		id, _ := uuid.NewV4()
		if result := db.Create(&databasemodels.IisApplicationPool{
			IisApplicationPoolId: id,
			Name:                 iisApp.ApplicationPool.Name,
			ProcessType:          iisApp.ApplicationPool.ProcessType,
			FrameworkVersion:     iisApp.ApplicationPool.FrameworkVersion,
		}); result.Error != nil {
			panic(result.Error)
		}

		id, _ = uuid.NewV4()
		var appPool databasemodels.IisApplicationPool
		if result := db.First(&appPool,
			&databasemodels.IisApplicationPool{Name: iisApp.ApplicationPool.Name}); result.Error != nil {
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
		if result := db.Create(&databasemodels.IisApplicationPool{
			IisApplicationPoolId: id,
			Name:                 iisSite.ApplicationPool.Name,
			ProcessType:          iisSite.ApplicationPool.ProcessType,
			FrameworkVersion:     iisSite.ApplicationPool.FrameworkVersion,
		}); result.Error != nil {
			panic(result.Error)
		}

		id, _ = uuid.NewV4()
		var appPool databasemodels.IisApplicationPool
		if result := db.First(&appPool,
			&databasemodels.IisApplicationPool{Name: iisSite.ApplicationPool.Name}); result.Error != nil {
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
		for _, name := range scheduledTask.Names {
			id, _ := uuid.NewV4()
			if result := db.Create(&databasemodels.ScheduledTask{
				ScheduledTaskId:           id,
				Name:                      name,
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
	}

	for _, environment := range i.Environments {
		id, _ := uuid.NewV4()
		var ids []string

		for _, server := range environment.Servers {
			id, _ := uuid.NewV4()
			ids = append(ids, id.String())
			if result := db.Create(&databasemodels.Server{
				ServerId: id,
				Name:     server.Name,
				Type:     server.Type,
			}); result.Error != nil {
				panic(result.Error)
			}
		}

		if result := db.Create(&databasemodels.Environment{
			EnvironmentId: id,
			Name:          environment.Name,
			BusinessLine:  environment.BusinessLine,
			Servers:       ids,
		}); result.Error != nil {
			panic(result.Error)
		}
	}

	for _, application := range i.Applications {
		appNames := application.Resources.IisApplications
		siteNames := application.Resources.IisSites
		taskNames := application.Resources.ScheduledTasks
		serviceNames := application.Resources.WindowsServices

		var iisApplications []databasemodels.IisApplication
		if result := db.Model(&databasemodels.IisApplication{}).
			Where("application_name in (?)", appNames).
			Find(&iisApplications); result.Error != nil {
			panic(result.Error)
		}

		var appIds []string
		linq.From(iisApplications).SelectT(func(application databasemodels.IisApplication) string {
			return application.IisApplicationId.String()
		}).ToSlice(&appIds)

		var iisSites []databasemodels.IisSite
		if result := db.Model(&databasemodels.IisSite{}).
			Where("site_name in (?)", siteNames).
			Find(&iisSites); result.Error != nil {
			panic(result.Error)
		}

		var siteIds []string
		linq.From(iisApplications).SelectT(func(site databasemodels.IisSite) string {
			return site.IisSiteId.String()
		}).ToSlice(&siteIds)

		var scheduledTasks []databasemodels.ScheduledTask
		if result := db.Model(&databasemodels.IisSite{}).
			Where("name in (?)", taskNames).
			Find(&scheduledTasks); result.Error != nil {
			panic(result.Error)
		}

		var taskIds []string
		linq.From(scheduledTasks).SelectT(func(task databasemodels.ScheduledTask) string {
			return task.ScheduledTaskId.String()
		}).ToSlice(&taskIds)

		var windowsServices []databasemodels.WindowsService
		if result := db.Model(&databasemodels.WindowsService{}).
			Where("name in (?)", serviceNames).
			Find(&windowsServices); result.Error != nil {
			panic(result.Error)
		}

		var serviceIds []string
		linq.From(windowsServices).SelectT(func(service databasemodels.WindowsService) string {
			return service.WindowsServiceId.String()
		}).ToSlice(&serviceIds)

		id, _ := uuid.NewV4()
		db.Create(databasemodels.Resource{
			DeliverableId:   id,
			RepositoryName:  application.Repository,
			SolutionName:    application.Solution,
			ProjectName:     application.Project,
			IisApplications: appIds,
			IisSites:        siteIds,
			ScheduledTasks:  taskIds,
			WindowsServices: serviceIds,
		})
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
