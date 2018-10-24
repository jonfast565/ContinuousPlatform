CREATE PROCEDURE [dbo].[ClearKeyValueCache] @MachineName NVARCHAR(MAX) = NULL
	,@CurrentMachineOnly BIT = 1
AS
BEGIN
    IF @MachineName IS NULL
    BEGIN
	    IF @CurrentMachineOnly = 1
		    DELETE
		    FROM dbo.KeyValueCache
		    WHERE MachineName = HOST_NAME();
	    ELSE
		    DELETE
		    FROM dbo.KeyValueCache
		    WHERE 1 = 1;
    END
    ELSE
    BEGIN
	    DELETE
	    FROM dbo.KeyValueCache
	    WHERE MachineName = @MachineName;
    END
END
