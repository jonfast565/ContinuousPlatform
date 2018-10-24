CREATE TABLE [dbo].[DeploymentLocations]
(
    [DeploymentLocationId] UNIQUEIDENTIFIER NOT NULL,
    [FriendlyName] NVARCHAR(255) NOT NULL,
    [PhysicalPath] NVARCHAR(2550) NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_DeploymentLocationId] PRIMARY KEY ([DeploymentLocationId]),
    CONSTRAINT [UC_DeploymentLocations] UNIQUE ([FriendlyName])
)
