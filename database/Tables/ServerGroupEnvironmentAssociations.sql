CREATE TABLE [dbo].[ServerGroupEnvironmentAssociations]
(
    [EnvironmentId] UNIQUEIDENTIFIER NOT NULL,
    [ServerGroupId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_ServerGroupEnvironmentAssociations] PRIMARY KEY ([EnvironmentId], [ServerGroupId]),
    CONSTRAINT [FK_ServerGroupEnvironmentAssociations_Environments] FOREIGN KEY ([EnvironmentId]) REFERENCES [Environments]([EnvironmentId]),
    CONSTRAINT [FK_ServerGroupEnvironmentAssociations_ServerGroups] FOREIGN KEY ([ServerGroupId]) REFERENCES [ServerGroups]([ServerGroupId])
)
