CREATE TABLE [dbo].[ResourceAssociations]
(
    [ResourceAssociationId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [DeliverableId] UNIQUEIDENTIFIER NOT NULL,
    [IisSiteId] UNIQUEIDENTIFIER NULL,
    [IisApplicationId] UNIQUEIDENTIFIER NULL,
    [WindowsServiceId] UNIQUEIDENTIFIER NULL,
    [ScheduledTaskId] UNIQUEIDENTIFIER NULL,
    [DeploymentLocationId] UNIQUEIDENTIFIER NULL,
    [ResourceAmalgamationId] UNIQUEIDENTIFIER NOT NULL,
    [Enabled] BIT NOT NULL DEFAULT 1,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [CH_ResourceAssociations] CHECK -- only one
    ((CASE WHEN [IisSiteId] IS NOT NULL THEN 1 ELSE 0 END
    + CASE WHEN [IisApplicationId] IS NOT NULL THEN 1 ELSE 0 END
    + CASE WHEN [WindowsServiceId] IS NOT NULL THEN 1 ELSE 0 END
    + CASE WHEN [ScheduledTaskId] IS NOT NULL THEN 1 ELSE 0 END
    + CASE WHEN [DeploymentLocationId] IS NOT NULL THEN 1 ELSE 0 END) = 1),
    CONSTRAINT [PK_ResourceAssociations] PRIMARY KEY ([ResourceAssociationId]),
    CONSTRAINT [FK_ResourceAssociations_Deliverables] FOREIGN KEY ([DeliverableId]) REFERENCES [Deliverables]([DeliverableId]),
    CONSTRAINT [FK_ResourceAssociations_IisSites] FOREIGN KEY ([IisSiteId]) REFERENCES [IisSites]([IisSiteId]),
    CONSTRAINT [FK_ResourceAssociations_IisApplications] FOREIGN KEY ([IisApplicationId]) REFERENCES [IisApplications]([IisApplicationId]),
    CONSTRAINT [FK_ResourceAssociations_WindowsServices] FOREIGN KEY ([WindowsServiceId]) REFERENCES [WindowsServices]([WindowsServiceId]),
    CONSTRAINT [FK_ResourceAssociations_ScheduledTasks] FOREIGN KEY ([ScheduledTaskId]) REFERENCES [ScheduledTasks]([ScheduledTaskId]),
    CONSTRAINT [FK_ResourceAssociations_DeploymentLocations] FOREIGN KEY ([DeploymentLocationId]) REFERENCES [DeploymentLocations]([DeploymentLocationId]),
    CONSTRAINT [FK_ResourceAssociations_ResourceAmalgamations] FOREIGN KEY ([ResourceAmalgamationId]) REFERENCES [ResourceAmalgamations]([ResourceAmalgamationId])
)
