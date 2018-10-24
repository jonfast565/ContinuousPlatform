CREATE TABLE [dbo].[ResourceAmalgamations]
(
    [ResourceAmalgamationId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [BusinessLineId] UNIQUEIDENTIFIER NOT NULL,
    [Name] NVARCHAR(255) NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [PK_ResourceAmalgamations] PRIMARY KEY ([ResourceAmalgamationId]),
    CONSTRAINT [UC_ResourceAmalgamations] UNIQUE ([Name]),
    CONSTRAINT [FK_ResourceAmalgamations_BusinessLines] FOREIGN KEY ([BusinessLineId]) REFERENCES [BusinessLines]([BusinessLineId])
)
