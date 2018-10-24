CREATE PROCEDURE GetApplicationLookupReport
AS
BEGIN
	SELECT DISTINCT e.EnvironmentName
		/*,(
			-- server info
			-- WARNING: REQUIRES SQL SERVER 2017!
			SELECT STRING_AGG(s.GroupName, ', ')
			FROM (
				SELECT DISTINCT sg.GroupName
				FROM [Servers] s
				INNER JOIN ServerTypes st ON st.ServerTypeId = s.ServerTypeId
				INNER JOIN ServerGroupAssociations sga ON s.ServerId = sga.ServerId
				INNER JOIN ServerGroups sg ON sga.ServerGroupId = sg.ServerGroupId
                INNER JOIN ServerGroupEnvironmentAssociations sgea ON sgea.ServerGroupId = sg.ServerGroupId
					AND sgea.EnvironmentId = e.EnvironmentId
				WHERE (
						CASE 
							WHEN ras.IisApplicationId IS NOT NULL
								OR ras.IisSiteId IS NOT NULL
								THEN 1
							ELSE 0
							END
						) = (
						CASE 
							WHEN st.ServerTypeName = 'Web'
								THEN 1
							ELSE 0
							END
						)
				) AS s
			) AS ServerGroupList */
		,(
			-- WARNING: REQUIRES SQL SERVER 2017!
			SELECT STRING_AGG(s.ServerName, ', ')
			FROM (
				SELECT DISTINCT ServerName
				FROM [Servers] s
				INNER JOIN ServerTypes st ON st.ServerTypeId = s.ServerTypeId
				INNER JOIN ServerGroupAssociations sga ON s.ServerId = sga.ServerId
				INNER JOIN ServerGroups sg ON sga.ServerGroupId = sg.ServerGroupId
				INNER JOIN ServerGroupEnvironmentAssociations sgea ON sgea.ServerGroupId = sg.ServerGroupId
					AND sgea.EnvironmentId = e.EnvironmentId
				WHERE (
						CASE 
							WHEN ras.IisApplicationId IS NOT NULL
								OR ras.IisSiteId IS NOT NULL
								THEN 1
							ELSE 0
							END
						) = (
						CASE 
							WHEN st.ServerTypeName = 'Web'
								THEN 1
							ELSE 0
							END
						)
				) AS s
			) AS ServerList
        -- deliverable info
		,d.RepositoryName AS RepositoryName
		,d.SolutionName AS SolutionName
		,d.ProjectName AS ProjectName
        -- iis application info
		,ia.ApplicationName AS IisApplicationName
		,iap.PoolName AS IisApplicationAppPoolName
		,REPLACE(ia.PhysicalPath, '{{SiteName}}', iss.SiteName) AS IisApplicationPhysicalPath
		,iss.SiteName AS IisApplicationSiteName
		,iaps1.PoolName AS IisApplicationSiteApplicationPool
		,iss.PhysicalPath AS IisApplicationSitePhysicalPath
        -- ,isg.GroupName AS IisSiteGroupName
        -- iis site application info (i.e. code deployed to metadata.ncdr.com)
		,iss2.SiteName AS IisSiteApplicationName
        ,iaps2.PoolName AS IisSiteApplicationAppPool
		,iss2.PhysicalPath AS IisSiteApplicationPhysicalPath
        -- ,isg2.GroupName AS IisSiteGroupName
        -- task info
		,st.TaskName AS ScheduledTaskName
		,st.BinaryPath AS ScheduledTaskBinaryPath
        ,st.BinaryExecutableName AS ScheduledTaskBinaryExecutableName
        ,st.BinaryExecutableArguments AS ScheduledTaskBinaryExecutableArguments
		,st.RepeatInterval AS ScheduledTaskRepeatInterval
		-- ,stg.TaskGroupName AS ScheduledTaskGroupName
        -- service info
		,ws.ServiceName AS WindowServiceName
		,ws.BinaryPath AS WindowsServiceBinaryPath
        ,ws.BinaryExecutableName AS WindowsServiceBinaryExecutableName
        ,ws.BinaryExecutableArguments AS WindowsServiceBinaryExecutableArguments
		-- ,wsg.GroupName AS WindowsServiceGroupName
	FROM BusinessLines bl
	-- all environments
	CROSS JOIN Environments e
	INNER JOIN ResourceAmalgamations ra ON bl.BusinessLineId = ra.BusinessLineId
	INNER JOIN ResourceAssociations ras ON ra.ResourceAmalgamationId = ras.ResourceAmalgamationId
	INNER JOIN Deliverables d ON d.DeliverableId = ras.DeliverableId
	-- iis application infos
	LEFT JOIN IisApplications ia ON ras.IisApplicationId = ia.IisApplicationId
	LEFT JOIN IisApplicationPools iap ON ia.IisApplicationPoolId = iap.IisApplicationPoolId
	LEFT JOIN IisApplicationGroupAssociations iaga ON iaga.IisApplicationId = ia.IisApplicationId
	LEFT JOIN IisApplicationGroups iag ON iaga.IisApplicationGroupId = iag.IisApplicationGroupId
	LEFT JOIN IisSiteApplicationGroupAssociations isags ON isags.IisApplicationGroupId = iaga.IisApplicationGroupId
	LEFT JOIN IisSites iss ON iss.IisSiteId = isags.IisSiteId
	LEFT JOIN IisApplicationPools iaps1 ON iss.IisSiteApplicationPoolId = iaps1.IisApplicationPoolId
	-- iis site information
	LEFT JOIN IisSiteGroupAssociations isga ON isga.IisSiteId = iss.IisSiteId
	LEFT JOIN IisSiteGroups isg ON isg.IisSiteGroupId = isga.IisSiteGroupId
		AND e.EnvironmentId = isg.EnvironmentId
	-- iis site information for site specific applications
	LEFT JOIN IisSites iss2 ON ras.IisSiteId = iss2.IisSiteId
	LEFT JOIN IisApplicationPools iaps2 ON iss2.IisSiteApplicationPoolId = iaps2.IisApplicationPoolId
	LEFT JOIN IisSiteGroupAssociations isga2 ON isga2.IisSiteId = iss2.IisSiteId
	LEFT JOIN IisSiteGroups isg2 ON isg2.IisSiteGroupId = isga2.IisSiteGroupId
		AND e.EnvironmentId = isg2.EnvironmentId
	-- scheduled task information
	LEFT JOIN ScheduledTasks st ON st.ScheduledTaskId = ras.ScheduledTaskId
	LEFT JOIN ScheduledTaskGroupAssociations stga ON stga.ScheduledTaskId = st.ScheduledTaskId
	LEFT JOIN ScheduledTaskGroups stg ON stg.ScheduledTaskGroupId = stga.ScheduledTaskGroupId
		AND stg.EnvironmentId = e.EnvironmentId
	-- windows service information
	LEFT JOIN WindowsServices ws ON ws.WindowsServiceId = ras.WindowsServiceId
	LEFT JOIN WindowsServiceGroupAssociations wsga ON wsga.WindowsServiceId = ws.WindowsServiceId
	LEFT JOIN WindowsServiceGroups wsg ON wsg.WindowsServiceGroupId = wsga.WindowsServiceGroupId
		AND wsg.EnvironmentId = e.EnvironmentId
	WHERE bl.[Name] = 'Ncdr'
		-- and (@EnvironmentId IS NULL OR e.EnvironmentId = @EnvironmentId)
		AND (
			CASE 
				WHEN isg.GroupName IS NOT NULL
					THEN 1
				ELSE 0
				END + CASE 
				WHEN isg2.GroupName IS NOT NULL
					THEN 1
				ELSE 0
				END + CASE 
				WHEN stg.TaskGroupName IS NOT NULL
					THEN 1
				ELSE 0
				END + CASE 
				WHEN wsg.GroupName IS NOT NULL
					THEN 1
				ELSE 0
				END = 1
			)
	ORDER BY EnvironmentName ASC
		,RepositoryName ASC
		,SolutionName ASC
		,ProjectName ASC
		,IisApplicationName ASC
		,IisApplicationSiteName ASC
		,IisApplicationSitePhysicalPath ASC
        ,IisSiteApplicationName ASC
        ,IisSiteApplicationPhysicalPath ASC
		,WindowServiceName ASC
		,WindowsServiceBinaryPath ASC
		,ScheduledTaskName ASC
		,ScheduledTaskBinaryPath ASC
		,ServerList ASC;

	RETURN 0;
END;
