CREATE TABLE [dbo].[IisSiteGroups]
(
    [IisSiteGroupId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [GroupName] NVARCHAR(255) NOT NULL,
    [EnvironmentId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_IisSiteGroups] PRIMARY KEY ([IisSiteGroupId]),
    CONSTRAINT [UC_IisSiteGroups] UNIQUE ([GroupName]),
    CONSTRAINT [FK_IisSiteGroups_Environments] FOREIGN KEY ([EnvironmentId]) REFERENCES [Environments]([EnvironmentId])
)
