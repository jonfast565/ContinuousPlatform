package server

import (
	"../../databasemodels"
	"../../models/inframodel"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func getResource(key inframodel.ResourceKey, db *gorm.DB) *databasemodels.Resource {
	var resource databasemodels.Resource
	if db.First(&resource, &databasemodels.Resource{
		RepositoryName: key.RepositoryName,
		SolutionName:   key.SolutionName,
		ProjectName:    key.ProjectName,
	}).RecordNotFound() {
		panic("resource not found")
	}
	return &resource
}

func getIisApplications(applicationGuids []string, db *gorm.DB) *[]inframodel.IisApplication {
	var applications []inframodel.IisApplication
	for _, iisApp := range applicationGuids {
		var iisApplication databasemodels.IisApplication
		var iisApplicationPool databasemodels.IisApplicationPool

		uid, _ := uuid.FromString(iisApp)
		if db.First(&iisApplication, databasemodels.IisApplication{
			IisApplicationId: uid,
		}).RecordNotFound() {
			panic("app not found")
		}

		uid = iisApplication.ApplicationPoolId
		if db.First(&iisApplicationPool, databasemodels.IisApplicationPool{
			IisApplicationPoolId: uid,
		}).RecordNotFound() {
			panic("app pool not found")
		}

		applications = append(applications, inframodel.IisApplication{
			ApplicationName: iisApplication.ApplicationName,
			PhysicalPath:    iisApplication.PhysicalPath,
			AppPool: inframodel.IisApplicationPool{
				AppPoolName:             iisApplicationPool.Name,
				AppPoolType:             iisApplicationPool.ProcessType,
				AppPoolFrameworkVersion: iisApplicationPool.FrameworkVersion,
				AppPoolGuid:             iisApplicationPool.IisApplicationPoolId.String(),
			},
			ApplicationGuid: iisApplication.IisApplicationId.String(),
		})
	}
	return &applications
}

func getIisSitesForEnvironment(
	environment *inframodel.Environment,
	allSites *[]inframodel.IisSite,
	application *inframodel.IisApplication) *[]inframodel.IisSitePart {
	var results []inframodel.IisSitePart
	for _, site := range *allSites {
		for _, siteEnvironment := range site.Environments {
			applicationFound := false
			for _, app := range site.Applications {
				if app.ApplicationName == application.ApplicationName {
					applicationFound = true
				}
			}
			if siteEnvironment.Name == environment.Name &&
				siteEnvironment.BusinessLine == environment.BusinessLine &&
				applicationFound {
				results = append(results, inframodel.IisSitePart{
					SiteName:     site.SiteName,
					PhysicalPath: site.PhysicalPath,
					SiteGuid:     site.SiteGuid,
					Environments: site.Environments,
				})
			}
		}
	}
	return &results
}

func getAllIisSites(db *gorm.DB) *[]inframodel.IisSite {
	var iisSites []databasemodels.IisSite
	db.Find(&iisSites, &databasemodels.IisSite{})
	var siteGuids []string
	for _, site := range iisSites {
		siteGuids = append(siteGuids, site.IisSiteId.String())
	}
	result := getIisSites(siteGuids, db)
	return result
}

func getIisSites(siteGuids []string, db *gorm.DB) *[]inframodel.IisSite {
	var sites []inframodel.IisSite
	for _, iisSiteUid := range siteGuids {
		var iisSite databasemodels.IisSite
		var iisApplicationPool databasemodels.IisApplicationPool

		uid, _ := uuid.FromString(iisSiteUid)
		if db.First(&iisSite, databasemodels.IisSite{
			IisSiteId: uid,
		}).RecordNotFound() {
			panic("site not found")
		}

		uid = iisSite.ApplicationPoolId
		if db.First(&iisApplicationPool, databasemodels.IisApplicationPool{
			IisApplicationPoolId: uid,
		}).RecordNotFound() {
			panic("app pool not found")
		}

		iisApplications := getIisApplications(iisSite.SiteApplications, db)
		environmentParts := getEnvironmentPartsByIds(iisSite.Environments, db)

		sites = append(sites, inframodel.IisSite{
			SiteName:     iisSite.SiteName,
			PhysicalPath: iisSite.PhysicalPath,
			AppPool: inframodel.IisApplicationPool{
				AppPoolName:             iisApplicationPool.Name,
				AppPoolType:             iisApplicationPool.ProcessType,
				AppPoolFrameworkVersion: iisApplicationPool.FrameworkVersion,
				AppPoolGuid:             iisApplicationPool.IisApplicationPoolId.String(),
			},
			SiteGuid:     iisSite.IisSiteId.String(),
			Applications: *iisApplications,
			Environments: environmentParts,
		})
	}
	return &sites
}

func getAllIisSiteParts(db *gorm.DB) *[]inframodel.IisSitePart {
	var iisSiteList []databasemodels.IisSite

	if db.Find(&iisSiteList, &databasemodels.IisSite{}).RecordNotFound() {
		panic("sites not found")
	}

	var results []inframodel.IisSitePart
	for _, iisSite := range iisSiteList {
		environmentParts := getEnvironmentPartsByIds(iisSite.Environments, db)
		results = append(results, inframodel.IisSitePart{
			SiteName:     iisSite.SiteName,
			PhysicalPath: iisSite.PhysicalPath,
			SiteGuid:     iisSite.IisSiteId.String(),
			Environments: environmentParts,
		})
	}
	return &results
}

func getRelevantAppPools(siteModels []inframodel.IisSite,
	appModels []inframodel.IisApplication) []inframodel.IisApplicationPool {
	if len(appModels) > 0 {
		var pools []inframodel.IisApplicationPool
		for _, appModel := range appModels {
			pools = append(pools, appModel.AppPool)
		}
		return pools
	}
	if len(siteModels) > 0 {
		var pools []inframodel.IisApplicationPool
		for _, siteModel := range siteModels {
			pools = append(pools, siteModel.AppPool)
		}
		return pools
	}
	return []inframodel.IisApplicationPool{}
}

func getAppPoolNames(models []inframodel.IisApplicationPool) []string {
	names := make([]string, 0)
	for _, model := range models {
		names = append(names, model.AppPoolName)
	}
	return names
}

func getDeploymentLocations(
	iisApps []inframodel.IisApplication,
	iisSites []inframodel.IisSite,
	scheduledTasks []inframodel.ScheduledTask,
	services []inframodel.WindowsService) []string {
	if len(iisApps) > 0 {
		var paths []string
		for _, iisApp := range iisApps {
			paths = append(paths, iisApp.PhysicalPath)
		}
		return paths
	}
	if len(iisSites) > 0 {
		var paths []string
		for _, iisSite := range iisSites {
			paths = append(paths, iisSite.PhysicalPath)
		}
		return paths
	}
	if len(scheduledTasks) > 0 {
		var paths []string
		for _, scheduledTask := range scheduledTasks {
			paths = append(paths, scheduledTask.BinaryPath)
		}
		return paths
	}
	if len(services) > 0 {
		var paths []string
		for _, service := range services {
			paths = append(paths, service.BinaryPath)
		}
		return paths
	}
	return []string{}
}

func getWindowsTasks(taskGuids []string, db *gorm.DB) *[]inframodel.ScheduledTask {
	var scheduledTasks []inframodel.ScheduledTask
	for _, scheduledTaskUid := range taskGuids {
		var scheduledTask databasemodels.WindowsScheduledTask

		uid, _ := uuid.FromString(scheduledTaskUid)
		if db.First(&scheduledTask, databasemodels.WindowsScheduledTask{
			WindowsScheduledTaskId: uid,
		}).RecordNotFound() {
			panic("scheduled task not found")
		}

		environmentParts := getEnvironmentPartsByIds(scheduledTask.Environments, db)
		scheduledTasks = append(scheduledTasks, inframodel.ScheduledTask{
			TaskName:                  scheduledTask.TaskName,
			BinaryPath:                scheduledTask.BinaryPath,
			BinaryExecutableName:      scheduledTask.BinaryExecutableName,
			BinaryExecutableArguments: scheduledTask.BinaryExecutableArguments,
			ScheduleType:              scheduledTask.ScheduleType,
			RepeatInterval:            scheduledTask.RepeatInterval,
			RepetitionDuration:        scheduledTask.RepetitionDuration,
			ExecutionTimeLimit:        scheduledTask.ExecutionTimeLimit,
			Priority:                  scheduledTask.Priority,
			TaskGuid:                  scheduledTask.WindowsScheduledTaskId.String(),
			Environments:              environmentParts,
		})
	}
	return &scheduledTasks
}

func getScheduledTaskNames(models []inframodel.ScheduledTask) []string {
	names := make([]string, 0)
	for _, model := range models {
		names = append(names, model.TaskName)
	}
	return names
}

func getWindowsServices(serviceGuids []string, db *gorm.DB) *[]inframodel.WindowsService {
	var services []inframodel.WindowsService
	for _, serviceUid := range serviceGuids {
		var service databasemodels.WindowsService

		uid, _ := uuid.FromString(serviceUid)
		if db.First(&service, databasemodels.WindowsService{
			WindowsServiceId: uid,
		}).RecordNotFound() {
			panic("windows service not found")
		}

		environmentParts := getEnvironmentPartsByIds(service.Environments, db)
		services = append(services, inframodel.WindowsService{
			ServiceName:               service.ServiceName,
			BinaryPath:                service.BinaryPath,
			BinaryExecutableName:      service.BinaryExecutableName,
			BinaryExecutableArguments: service.BinaryExecutableArguments,
			ServiceGuid:               service.WindowsServiceId.String(),
			Environments:              environmentParts,
		})
	}
	return &services
}

func getWindowsServiceNames(models []inframodel.WindowsService) []string {
	names := make([]string, 0)
	for _, model := range models {
		names = append(names, model.ServiceName)
	}
	return names
}

func getEnvironmentPartsByIds(environmentIds []string, db *gorm.DB) []inframodel.EnvironmentPart {
	var databaseEnvironments []databasemodels.Environment
	if result := db.Model(&databasemodels.Environment{}).
		Where("environment_id in (?)", environmentIds).
		Find(&databaseEnvironments); result.Error != nil {
		panic(result.Error)
	}
	var environmentParts []inframodel.EnvironmentPart
	for _, databaseEnvironment := range databaseEnvironments {
		environmentParts = append(environmentParts, inframodel.EnvironmentPart{
			BusinessLine: databaseEnvironment.BusinessLine,
			Name:         databaseEnvironment.Name,
		})
	}
	return environmentParts
}

func getEnvironments(db *gorm.DB) *inframodel.EnvironmentList {
	var environments []databasemodels.Environment
	if db.Find(&environments).RecordNotFound() {
		panic("environments empty")
	}

	var environmentResult inframodel.EnvironmentList
	for _, environment := range environments {
		serverUids := environment.Servers
		serverUidStrArray := []string(serverUids)
		var servers []databasemodels.Server
		if result := db.Model(&databasemodels.Server{}).
			Where("server_id in (?)", serverUidStrArray).
			Find(&servers); result.Error != nil {
			panic(result.Error)
		}

		var infraServers []inframodel.ServerType
		for _, server := range servers {
			infraServers = append(infraServers, inframodel.ServerType{
				ServerName: server.Name,
			})
		}

		environmentResult = append(environmentResult, inframodel.Environment{
			BusinessLine: environment.BusinessLine,
			Name:         environment.Name,
			Servers:      infraServers,
		})
	}
	return &environmentResult
}

func filterEnvironments(
	resourceIisApplications []inframodel.IisApplication,
	allIisSites []inframodel.IisSite,
	resourceIisSites []inframodel.IisSite,
	resourceScheduledTask []inframodel.ScheduledTask,
	resourceServices []inframodel.WindowsService,
	allEnvironments inframodel.EnvironmentList) *inframodel.EnvironmentList {
	if len(resourceIisApplications) > 0 {
		envs := getIisApplicationEnvironments(allIisSites, resourceIisApplications, allEnvironments)
		return &envs
	}
	if len(resourceIisSites) > 0 {
		envs := getSiteEnvironments(allEnvironments, resourceIisSites)
		return &envs
	}
	if len(resourceScheduledTask) > 0 {
		envs := getTaskEnvironments(allEnvironments, resourceScheduledTask)
		return &envs
	}
	if len(resourceServices) > 0 {
		envs := getServiceEnvironments(allEnvironments, resourceServices)
		return &envs
	}
	return &inframodel.EnvironmentList{}
}

func getSiteEnvironments(
	allEnvironments []inframodel.Environment,
	resourceIisSites []inframodel.IisSite) inframodel.EnvironmentList {
	var environmentResultList inframodel.EnvironmentList
	for _, iisSite := range resourceIisSites {
		for _, siteEnvironment := range iisSite.Environments {
			for _, environment := range allEnvironments {
				if siteEnvironment.NameMatchEnv(environment) &&
					!environmentResultList.HasMatch(environment) {
					environmentResultList = append(environmentResultList, environment)
				}
			}
		}
	}
	return environmentResultList
}

func getIisApplicationEnvironments(
	allIisSites []inframodel.IisSite,
	apps []inframodel.IisApplication,
	allEnvironments []inframodel.Environment) inframodel.EnvironmentList {
	var environmentResultList inframodel.EnvironmentList
	for _, iisSite := range allIisSites {
		for _, siteApplication := range iisSite.Applications {
			for _, resourceApplication := range apps {
				if siteApplication.ApplicationGuid == resourceApplication.ApplicationGuid {
					for _, siteEnvironment := range iisSite.Environments {
						for _, environment := range allEnvironments {
							if siteEnvironment.NameMatchEnv(environment) &&
								!environmentResultList.HasMatch(environment) {
								environmentResultList = append(environmentResultList, environment)
							}
						}
					}
				}
			}
		}
	}
	return environmentResultList
}

func getTaskEnvironments(
	allEnvironments []inframodel.Environment,
	resourceScheduledTask []inframodel.ScheduledTask) inframodel.EnvironmentList {
	var environmentResultList inframodel.EnvironmentList
	for _, environment := range allEnvironments {
		for _, task := range resourceScheduledTask {
			for _, taskEnvironment := range task.Environments {
				if taskEnvironment.NameMatchEnv(environment) &&
					!environmentResultList.HasMatch(environment) {
					environmentResultList = append(environmentResultList, environment)
				}
			}
		}
	}
	return environmentResultList
}

func getServiceEnvironments(
	allEnvironments inframodel.EnvironmentList,
	resourceServices []inframodel.WindowsService) inframodel.EnvironmentList {
	var environmentResultList inframodel.EnvironmentList
	for _, environment := range allEnvironments {
		for _, service := range resourceServices {
			for _, serviceEnvironment := range service.Environments {
				if serviceEnvironment.NameMatchEnv(environment) &&
					!environmentResultList.HasMatch(environment) {
					environmentResultList = append(environmentResultList, environment)
				}
			}
		}
	}
	return environmentResultList
}
