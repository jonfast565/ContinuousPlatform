CREATE TABLE [dbo].[IisApplications]
(
    [IisApplicationId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [IisApplicationPoolId] UNIQUEIDENTIFIER NOT NULL,
    [ApplicationName] NVARCHAR(255) NOT NULL,
    [PhysicalPath] NVARCHAR(2550) NOT NULL,
    [ApplicationInternalAliasName] NVARCHAR(255) NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_IisApplications] PRIMARY KEY ([IisApplicationId]),
    CONSTRAINT [FK_IisApplications_IisApplicationPools] FOREIGN KEY ([IisApplicationPoolId]) REFERENCES [IisApplicationPools]([IisApplicationPoolId])
)
