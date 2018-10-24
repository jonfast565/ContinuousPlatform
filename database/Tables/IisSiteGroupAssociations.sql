CREATE TABLE [dbo].[IisSiteGroupAssociations]
(
    [IisSiteGroupId] UNIQUEIDENTIFIER NOT NULL,
    [IisSiteId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_IisSiteGroupAssignments] PRIMARY KEY ([IisSiteGroupId], [IisSiteId]),
    CONSTRAINT [FK_IisSiteGroupAssignments_IisSiteGroups] FOREIGN KEY ([IisSiteGroupId]) REFERENCES [IisSiteGroups]([IisSiteGroupId]),
    CONSTRAINT [FK_IisSiteGroupAssignments_IisSites] FOREIGN KEY ([IisSiteId]) REFERENCES [IisSites]([IisSiteId])
)
