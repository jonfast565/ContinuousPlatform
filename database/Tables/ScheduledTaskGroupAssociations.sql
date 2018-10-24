CREATE TABLE [dbo].[ScheduledTaskGroupAssociations]
(
    [ScheduledTaskId] UNIQUEIDENTIFIER NOT NULL,
    [ScheduledTaskGroupId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_ScheduledTaskGroupAssignments] PRIMARY KEY ([ScheduledTaskId], [ScheduledTaskGroupId]),
    CONSTRAINT [FK_ScheduledTaskGroupAssignments_ScheduledTaskId] FOREIGN KEY ([ScheduledTaskId]) REFERENCES [ScheduledTasks]([ScheduledTaskId]),
    CONSTRAINT [FK_ScheduledTaskGroupAssignments_ScheduledTaskGroupId] FOREIGN KEY ([ScheduledTaskGroupId]) REFERENCES [ScheduledTaskGroups]([ScheduledTaskGroupId])
)
