CREATE TABLE [dbo].[ScheduledTaskGroups]
(
    [ScheduledTaskGroupId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [TaskGroupName] NVARCHAR(255) NOT NULL,
    [EnvironmentId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_ScheduledTaskGroups] PRIMARY KEY ([ScheduledTaskGroupId]),
    CONSTRAINT [UC_TaskGroupName] UNIQUE ([TaskGroupName]),
    CONSTRAINT [FK_ScheduledTaskGroups_Environments] FOREIGN KEY ([EnvironmentId]) REFERENCES [Environments]([EnvironmentId])
)
