CREATE TABLE [dbo].[DeploymentLocationGroups]
(
    [DeploymentLocationGroupId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [GroupName] NVARCHAR(255) NOT NULL,
    [EnvironmentId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_DeploymentLocationGroups] PRIMARY KEY ([DeploymentLocationGroupId]),
    CONSTRAINT [FK_DeploymentLocationGroups_Environments] FOREIGN KEY ([EnvironmentId]) REFERENCES [Environments]([EnvironmentId]),
    CONSTRAINT [UC_DeploymentLocationGroups] UNIQUE ([GroupName])
)
