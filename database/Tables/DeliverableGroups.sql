CREATE TABLE [dbo].[DeliverableGroups]
(
    [DeliverableGroupId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [GroupName] NVARCHAR(255) NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_DeliverableGroups] PRIMARY KEY ([DeliverableGroupId]),
    CONSTRAINT [UC_DeliverableGroups] UNIQUE ([GroupName])
)
