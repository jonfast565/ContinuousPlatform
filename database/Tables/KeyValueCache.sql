﻿CREATE TABLE [dbo].[KeyValueCache]
(
    [KeyValueCacheId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [Key] NVARCHAR(255) NOT NULL,
    [Value] VARBINARY(MAX) NOT NULL,
    [ValueType] NVARCHAR(2550) NOT NULL,
    [MachineName] NVARCHAR(255) NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_KeyValueCache] PRIMARY KEY ([KeyValueCacheId]),
    CONSTRAINT [UC_KeyValueCache] UNIQUE ([Key], [MachineName])
)