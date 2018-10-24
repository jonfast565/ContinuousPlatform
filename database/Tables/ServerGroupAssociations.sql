CREATE TABLE [dbo].[ServerGroupAssociations]
(
    [ServerId] UNIQUEIDENTIFIER NOT NULL,
    [ServerGroupId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [FK_ServerGroupAssociations_Servers] FOREIGN KEY ([ServerId]) REFERENCES [Servers]([ServerId]),
    CONSTRAINT [FK_ServerGroupAssociations_ServerGroups] FOREIGN KEY ([ServerGroupId]) REFERENCES [ServerGroups]([ServerGroupId]),
    CONSTRAINT [PK_ServerGroupAssociations] PRIMARY KEY ([ServerId], [ServerGroupId])
)
