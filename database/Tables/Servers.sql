CREATE TABLE [dbo].[Servers]
(
    [ServerId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [ServerName] NVARCHAR(255) NOT NULL,
    [ServerTypeId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_Servers] PRIMARY KEY ([ServerId]),
    CONSTRAINT [UC_Servers] UNIQUE ([ServerName]),
    CONSTRAINT [FK_Servers_ServerTypes] FOREIGN KEY ([ServerTypeId]) REFERENCES [ServerTypes]([ServerTypeId])
)
