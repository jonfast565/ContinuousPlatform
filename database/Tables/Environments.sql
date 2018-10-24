CREATE TABLE [dbo].[Environments]
(
    [EnvironmentId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [EnvironmentName] NVARCHAR(255) NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_Environments] PRIMARY KEY ([EnvironmentId]),
    CONSTRAINT [UC_Environments] UNIQUE ([EnvironmentName])
)
