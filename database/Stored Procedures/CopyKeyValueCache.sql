CREATE PROCEDURE [dbo].[CopyKeyValueCache] @SourceMachineName NVARCHAR(MAX)
	,@TargetMachineName NVARCHAR(MAX)
AS
IF @SourceMachineName IS NULL
	OR @TargetMachineName IS NULL THROW 51000
	,'Source machine or target machine not specified'
	,1;
	BEGIN TRANSACTION T1;

DELETE
FROM KeyValueCache
WHERE MachineName = @TargetMachineName;

INSERT INTO KeyValueCache (
	[Key]
	,[Value]
	,ValueType
	,MachineName
	)
SELECT [Key]
	,[Value]
	,ValueType
	,@TargetMachineName
FROM KeyValueCache
WHERE MachineName = @SourceMachineName;

COMMIT TRANSACTION T1;

RETURN 0
