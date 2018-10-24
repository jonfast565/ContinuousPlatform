CREATE TABLE [dbo].[DeliverableGroupAssociations]
(
    [DeliverableId] UNIQUEIDENTIFIER NOT NULL,
    [DeliverableGroupId] UNIQUEIDENTIFIER NOT NULL,
    [CreatedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [CreatedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    [LastModifiedBy] NVARCHAR(255) NOT NULL DEFAULT(SUSER_SNAME()),
    [LastModifiedDateTime] DATETIME NOT NULL DEFAULT(GETDATE()),
    CONSTRAINT [FK_DeliverableGroupAssociations_DeliverableGroups] FOREIGN KEY ([DeliverableGroupId]) REFERENCES [DeliverableGroups]([DeliverableGroupId]),
    CONSTRAINT [FK_DeliverableGroupAssociations_Deliverables] FOREIGN KEY ([DeliverableId]) REFERENCES [Deliverables]([DeliverableId]),
    CONSTRAINT [PK_DeliverableGroupAssociations] PRIMARY KEY([DeliverableId], [DeliverableGroupId])
)
