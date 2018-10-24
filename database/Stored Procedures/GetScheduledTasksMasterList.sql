CREATE PROCEDURE GetScheduledTasksMasterList
AS
BEGIN
	SELECT e.EnvironmentName
		-- ,stg.TaskGroupName AS TaskGroupName
		,st.TaskName AS ScheduledTaskName
		,st.BinaryPath
        ,st.BinaryExecutableName
        ,st.BinaryExecutableArguments
		,st.[Priority] AS [Priority]
		,st.ScheduleType
		,st.ExecutionTimeLimit
		,st.RepeatInterval
		,st.RepetitionDuration
	FROM Environments e
	INNER JOIN ScheduledTaskGroups stg ON stg.EnvironmentId = e.EnvironmentId
	INNER JOIN ScheduledTaskGroupAssociations stga ON stga.ScheduledTaskGroupId = stg.ScheduledTaskGroupId
	INNER JOIN ScheduledTasks st ON st.ScheduledTaskId = stga.ScheduledTaskId
	ORDER BY e.EnvironmentName ASC
		,stg.TaskGroupName ASC
		,st.TaskName ASC
		,st.BinaryPath ASC;

	RETURN 0;
END;
GO


