CREATE TABLE [dbo].[WindowsServices]
(
    [WindowsServiceId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [ServiceName] NVARCHAR(255) NOT NULL,
    [BinaryPath] NVARCHAR(2550) NOT NULL,
    [BinaryExecutableName] NVARCHAR(255) NOT NULL,
    [BinaryExecutableArguments] NVARCHAR(255) NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_WindowsServices] PRIMARY KEY ([WindowsServiceId]),
    CONSTRAINT [UC_WindowsServices] UNIQUE ([ServiceName])
)
