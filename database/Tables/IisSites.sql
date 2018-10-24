CREATE TABLE [dbo].[IisSites]
(
    [IisSiteId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [IisSiteApplicationPoolId] UNIQUEIDENTIFIER NOT NULL,
    [SiteName] NVARCHAR(255) NOT NULL,
    [PhysicalPath] NVARCHAR(2550) NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_IisSites] PRIMARY KEY ([IisSiteId]),
    CONSTRAINT [UC_IisSites] UNIQUE ([SiteName]),
    CONSTRAINT [FK_IisSites_IisApplicationPools] FOREIGN KEY ([IisSiteApplicationPoolId]) REFERENCES [IisApplicationPools]([IisApplicationPoolId])
)
