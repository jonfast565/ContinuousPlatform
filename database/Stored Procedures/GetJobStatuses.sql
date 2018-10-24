CREATE PROCEDURE [dbo].[GetJobStatuses]
    @MachineName NVARCHAR(255)
AS
BEGIN
    SELECT j.JobName,
           j.MachineName,
           j.JobTrigger,
           s.StateName AS JobState
    FROM dbo.Jobs j
        INNER JOIN dbo.JobStates s
            ON s.JobStateId = j.JobStateId
    WHERE j.MachineName = @MachineName
    ORDER BY j.JobOrder ASC;
END
