CREATE TABLE [dbo].[IisApplicationPools]
(
    [IisApplicationPoolId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [PoolName] NVARCHAR(255) NOT NULL,
    [PoolType] NVARCHAR(255) NOT NULL,
    [PoolFrameworkVersion] NVARCHAR(255) NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_IisApplicationPools] PRIMARY KEY ([IisApplicationPoolId]),
    CONSTRAINT [UC_IisApplicationPools] UNIQUE ([PoolName])
)
