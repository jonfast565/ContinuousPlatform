CREATE TABLE [dbo].[IisSiteBindings]
(
    [IisSiteBindingId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [IisSiteId] UNIQUEIDENTIFIER NOT NULL,
    [BindingName] NVARCHAR(255) NOT NULL,
    [Port] INT NOT NULL,
    [CertificateThumbprint] NVARCHAR(2550) NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_IisSiteBindings] PRIMARY KEY ([IisSiteBindingId]),
    CONSTRAINT [FK_IisSiteBindings_IisSites] FOREIGN KEY ([IisSiteId]) REFERENCES [IisSites]([IisSiteId])
)
