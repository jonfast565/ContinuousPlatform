CREATE TABLE [dbo].[ScheduledTasks]
(
    [ScheduledTaskId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [TaskName] NVARCHAR(255) NOT NULL,
    [BinaryPath] NVARCHAR(2550) NOT NULL,
    [BinaryExecutableName] NVARCHAR(255) NOT NULL,
    [BinaryExecutableArguments] NVARCHAR(255) NULL,
    [ScheduleType] NVARCHAR(255) NULL,
    [RepeatInterval] BIGINT NULL,
    [RepetitionDuration] BIGINT NULL,
    [ExecutionTimeLimit] BIGINT NULL,
    [Priority] INT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_ScheduledTasks] PRIMARY KEY ([ScheduledTaskId]),
    CONSTRAINT [UC_ScheduledTasks] UNIQUE ([TaskName])
)
