CREATE TABLE [dbo].[DeploymentLocationGroupAssociations]
(
    [DeploymentLocationId] UNIQUEIDENTIFIER NOT NULL,
    [DeploymentLocationGroupId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_DeploymentLocationGroupAssociations] PRIMARY KEY ([DeploymentLocationId], [DeploymentLocationGroupId]),
    CONSTRAINT [FK_DeploymentLocationGroupAssociations_DeploymentLocations] FOREIGN KEY ([DeploymentLocationId]) REFERENCES [DeploymentLocations]([DeploymentLocationId]),
    CONSTRAINT [FK_DeploymentLocationGroupAssociations_DeploymentLocationGroups] FOREIGN KEY ([DeploymentLocationGroupId]) REFERENCES [DeploymentLocationGroups]([DeploymentLocationGroupId]),
)
