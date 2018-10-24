-- PRE DEPLOYMENT SCRIPT

IF OBJECT_ID('KeyValueCache', 'U') IS NOT NULL
BEGIN
	SELECT *
	INTO #TempKeyValueCache
	FROM KeyValueCache;

	DELETE
	FROM KeyValueCache;
END;

IF OBJECT_ID('Logs', 'U') IS NOT NULL
    DELETE 
    FROM Logs;

IF OBJECT_ID('Jobs', 'U') IS NOT NULL
    DELETE
    FROM Jobs;

IF OBJECT_ID('JobStates', 'U') IS NOT NULL
    DELETE
    FROM JobStates;

IF OBJECT_ID('ResourceAssociations', 'U') IS NOT NULL
	DELETE
	FROM ResourceAssociations;

IF OBJECT_ID('ResourceAmalgamations', 'U') IS NOT NULL
	DELETE
	FROM ResourceAmalgamations;

IF OBJECT_ID('BusinessLines', 'U') IS NOT NULL
	DELETE
	FROM BusinessLines;

IF OBJECT_ID('DeploymentLocationGroupAssociations', 'U') IS NOT NULL
	DELETE
	FROM DeploymentLocationGroupAssociations;

IF OBJECT_ID('DeploymentLocations', 'U') IS NOT NULL
	DELETE
	FROM DeploymentLocations;

IF OBJECT_ID('DeploymentLocationGroups', 'U') IS NOT NULL
	DELETE
	FROM DeploymentLocationGroups;

IF OBJECT_ID('DeliverableGroups', 'U') IS NOT NULL
	DELETE
	FROM DeliverableGroupAssociations;

IF OBJECT_ID('Deliverables', 'U') IS NOT NULL
	DELETE
	FROM Deliverables;

IF OBJECT_ID('DeliverableGroups', 'U') IS NOT NULL
	DELETE
	FROM DeliverableGroups;

IF OBJECT_ID('ServerGroupEnvironmentAssociations', 'U') IS NOT NULL
	DELETE
	FROM ServerGroupEnvironmentAssociations;

IF OBJECT_ID('ServerGroupAssociations', 'U') IS NOT NULL
	DELETE
	FROM ServerGroupAssociations;

IF OBJECT_ID('ServerGroups', 'U') IS NOT NULL
	DELETE
	FROM ServerGroups;

IF OBJECT_ID('Servers', 'U') IS NOT NULL
	DELETE
	FROM Servers;

IF OBJECT_ID('ServerTypes', 'U') IS NOT NULL
	DELETE
	FROM ServerTypes;

IF OBJECT_ID('IisApplicationGroupAssociations', 'U') IS NOT NULL
	DELETE
	FROM IisApplicationGroupAssociations;

IF OBJECT_ID('IisSiteApplicationGroupAssociations', 'U') IS NOT NULL
	DELETE
	FROM IisSiteApplicationGroupAssociations;

IF OBJECT_ID('IisSiteGroupAssociations', 'U') IS NOT NULL
	DELETE
	FROM IisSiteGroupAssociations;

IF OBJECT_ID('IisApplications', 'U') IS NOT NULL
	DELETE
	FROM IisApplications;

IF OBJECT_ID('IisApplicationGroups', 'U') IS NOT NULL
	DELETE
	FROM IisApplicationGroups;

IF OBJECT_ID('IisSiteBindings', 'U') IS NOT NULL
	DELETE
	FROM IisSiteBindings;

IF OBJECT_ID('IisSites', 'U') IS NOT NULL
	DELETE
	FROM IisSites;

IF OBJECT_ID('IisSiteGroups', 'U') IS NOT NULL
	DELETE
	FROM IisSiteGroups;

IF OBJECT_ID('IisApplicationPools', 'U') IS NOT NULL
	DELETE
	FROM IisApplicationPools;

IF OBJECT_ID('ScheduledTaskGroupAssociations', 'U') IS NOT NULL
	DELETE
	FROM ScheduledTaskGroupAssociations;

IF OBJECT_ID('ScheduledTasks', 'U') IS NOT NULL
	DELETE
	FROM ScheduledTasks;

IF OBJECT_ID('ScheduledTaskGroups', 'U') IS NOT NULL
	DELETE
	FROM ScheduledTaskGroups;

IF OBJECT_ID('WindowsServiceGroupAssociations', 'U') IS NOT NULL
	DELETE
	FROM WindowsServiceGroupAssociations;

IF OBJECT_ID('WindowsServices', 'U') IS NOT NULL
	DELETE
	FROM WindowsServices;

IF OBJECT_ID('WindowsServiceGroups', 'U') IS NOT NULL
	DELETE
	FROM WindowsServiceGroups;

IF OBJECT_ID('Environments', 'U') IS NOT NULL
	DELETE
	FROM Environments;
