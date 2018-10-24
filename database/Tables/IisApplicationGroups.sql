CREATE TABLE [dbo].[IisApplicationGroups]
(
    [IisApplicationGroupId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [GroupName] NVARCHAR(255) NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_IisApplicationGroupId] PRIMARY KEY ([IisApplicationGroupId]),
    CONSTRAINT [UC_IisApplicationGroups] UNIQUE ([GroupName])
)
