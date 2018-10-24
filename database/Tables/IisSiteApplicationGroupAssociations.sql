CREATE TABLE [dbo].[IisSiteApplicationGroupAssociations]
(
    [IisSiteId] UNIQUEIDENTIFIER NOT NULL,
    [IisApplicationGroupId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_IisSiteApplicationGroupAssignments] PRIMARY KEY ([IisApplicationGroupId], [IisSiteId]),
    CONSTRAINT [FK_IisSiteApplicationGroupAssignments_IisSites] FOREIGN KEY ([IisSiteId]) REFERENCES [IisSites]([IisSiteId]),
    CONSTRAINT [FK_IisSiteApplicationGroupAssignments_IisApplicationGroups] FOREIGN KEY ([IisApplicationGroupId]) REFERENCES [IisApplicationGroups]([IisApplicationGroupId])
)
