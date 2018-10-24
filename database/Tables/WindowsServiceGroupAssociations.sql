CREATE TABLE [dbo].[WindowsServiceGroupAssociations]
(
    [WindowsServiceId] UNIQUEIDENTIFIER NOT NULL,
    [WindowsServiceGroupId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_WindowsServiceGroupAssignments] PRIMARY KEY ([WindowsServiceId], [WindowsServiceGroupId]),
    CONSTRAINT [FK_WindowsServiceGroupAssignments_WindowsServiceId] FOREIGN KEY ([WindowsServiceId]) REFERENCES [WindowsServices]([WindowsServiceId]),
    CONSTRAINT [FK_WindowsServiceGroupAssignments_WindowsServiceGroupId] FOREIGN KEY ([WindowsServiceGroupId]) REFERENCES [WindowsServiceGroups]([WindowsServiceGroupId])
)
