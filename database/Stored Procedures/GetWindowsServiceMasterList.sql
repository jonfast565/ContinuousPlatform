CREATE PROCEDURE GetWindowsServiceMasterList
AS
BEGIN
	SELECT e.EnvironmentName
		--,wsg.GroupName AS ServiceGroupName
		,ws.ServiceName AS ServiceName
		,ws.BinaryPath AS ServicePath
        ,ws.BinaryExecutableName
        ,ws.BinaryExecutableArguments
	FROM Environments e
	INNER JOIN WindowsServiceGroups wsg ON wsg.EnvironmentId = e.EnvironmentId
	INNER JOIN WindowsServiceGroupAssociations wsga ON wsga.WindowsServiceGroupId = wsg.WindowsServiceGroupId
	INNER JOIN WindowsServices ws ON ws.WindowsServiceId = wsga.WindowsServiceId
	ORDER BY e.EnvironmentName ASC
		,wsg.GroupName ASC
		,ws.ServiceName ASC
		,ws.BinaryPath ASC

	RETURN 0;
END;
GO


