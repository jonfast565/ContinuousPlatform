CREATE TABLE [dbo].[IisApplicationGroupAssociations]
(
    [IisApplicationGroupId] UNIQUEIDENTIFIER NOT NULL,
    [IisApplicationId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_IisApplicationGroupAssignments] PRIMARY KEY ([IisApplicationGroupId], [IisApplicationId]),
    CONSTRAINT [FK_IisApplicationGroupAssignments_IisApplicationGroups] FOREIGN KEY ([IisApplicationGroupId]) REFERENCES [IisApplicationGroups]([IisApplicationGroupId]),
    CONSTRAINT [FK_IisApplicationGroupAssignments_IisApplications] FOREIGN KEY ([IisApplicationId]) REFERENCES [IisApplications]([IisApplicationId])
)
