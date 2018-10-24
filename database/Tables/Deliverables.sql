CREATE TABLE [dbo].[Deliverables]
(
    [DeliverableId] UNIQUEIDENTIFIER NOT NULL DEFAULT NEWID(),
    [RepositoryName] NVARCHAR(255) NOT NULL,
    [SolutionName] NVARCHAR(255) NOT NULL,
    [ProjectName] NVARCHAR(255) NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE())
    CONSTRAINT [PK_Deliverables] PRIMARY KEY ([DeliverableId]),
    CONSTRAINT [UC_Deliverables] UNIQUE([RepositoryName], [SolutionName], [ProjectName])
)
