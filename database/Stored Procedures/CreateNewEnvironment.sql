CREATE PROCEDURE [dbo].[CreateNewEnvironment] @NewEnvironmentName NVARCHAR(MAX)
	,@ServerGroupName NVARCHAR(MAX)
	,@SiteGroupName NVARCHAR(MAX)
	,@SiteName NVARCHAR(MAX)
	,@SitePhysicalPath NVARCHAR(MAX)
	,@ApplicationGroupName NVARCHAR(MAX) = NULL
AS
BEGIN
	-- global errors
	DECLARE @ErrorMessage NVARCHAR(MAX)
	-- site constants
	DECLARE @SitePoolType NVARCHAR(MAX) = 'Integrated'
	DECLARE @SitePoolFrameworkVersion NVARCHAR(MAX) = 'v4.0'
	-- iis id tables
	DECLARE @IisSiteIds UniqueIdTable;
	DECLARE @IisSiteGroupIds UniqueIdTable;
	DECLARE @IisSiteAppPools UniqueIdTable;
	-- required variables
	DECLARE @EnvironmentId UNIQUEIDENTIFIER
	DECLARE @ServerGroupId UNIQUEIDENTIFIER

	-- parameter checks
	IF @NewEnvironmentName IS NULL
	BEGIN
		SET @ErrorMessage = '@ServerGroupName is NULL';

		THROW 51000
			,@ErrorMessage
			,1;
	END;

	IF (@ServerGroupName IS NULL)
	BEGIN
		SET @ErrorMessage = '@ServerGroupName is NULL';

		THROW 51000
			,@ErrorMessage
			,1;
	END;

	IF (@SiteGroupName IS NULL)
	BEGIN
		SET @ErrorMessage = '@ServerGroupName is NULL';

		THROW 51000
			,@ErrorMessage
			,1;
	END;

	IF (@SiteName IS NULL)
	BEGIN
		SET @ErrorMessage = '@SiteName is NULL';

		THROW 51000
			,@ErrorMessage
			,1;
	END;

	IF (@SitePhysicalPath IS NULL)
	BEGIN
		SET @ErrorMessage = '@SitePhysicalPath is NULL';

		THROW 51000
			,@ErrorMessage
			,1;
	END;

	SELECT @ServerGroupId = ServerGroupId
	FROM ServerGroups
	WHERE GroupName = @ServerGroupName;

	IF @ServerGroupId IS NULL
	BEGIN
		SELECT @ErrorMessage = CONCAT (
				'Server group '
				,@ServerGroupName
				,' does not exist'
				);

		THROW 51000
			,@ErrorMessage
			,1;
	END;

	BEGIN TRANSACTION CreateSites;

	INSERT INTO Environments (EnvironmentName)
	VALUES (@NewEnvironmentName);

	SELECT @EnvironmentId = EnvironmentId
	FROM Environments
	WHERE EnvironmentName = @NewEnvironmentName;

	INSERT INTO ServerGroupEnvironmentAssociations (
		ServerGroupId
		,EnvironmentId
		)
	VALUES (
		@ServerGroupId
		,@EnvironmentId
		);

	INSERT INTO IisApplicationPools (
		PoolName
		,PoolType
		,PoolFrameworkVersion
		)
	OUTPUT inserted.IisApplicationPoolId
	INTO @IisSiteAppPools([UniqueId])
	VALUES (
		@SiteName
		,@SitePoolType
		,@SitePoolFrameworkVersion
		);

	INSERT INTO IisSites (
		SiteName
		,PhysicalPath
		,IisSiteApplicationPoolId
		)
	OUTPUT inserted.IisSiteId
	INTO @IisSiteIds([UniqueId])
	VALUES (
		@SiteName
		,@SitePhysicalPath
		,(
			SELECT TOP 1 [UniqueId] AS IisSiteApplicationPoolId
			FROM @IisSiteAppPools
			)
		);

	INSERT INTO IisSiteGroups (
		GroupName
		,EnvironmentId
		)
	OUTPUT inserted.IisSiteGroupId
	INTO @IisSiteGroupIds([UniqueId])
	VALUES (
		@SiteGroupName
		,@EnvironmentId
		);

	INSERT INTO IisSiteGroupAssociations (
		IisSiteId
		,IisSiteGroupId
		)
	SELECT a.UniqueId AS IisSiteId
		,b.UniqueId AS IisSiteGroupId
	FROM @IisSiteIds a
	CROSS JOIN @IisSiteGroupIds b;

	IF @ApplicationGroupName IS NOT NULL
	BEGIN
		DECLARE @ApplicationGroupId UNIQUEIDENTIFIER

		SELECT @ApplicationGroupId = IisApplicationGroupId
		FROM IisApplicationGroups
		WHERE GroupName = @ApplicationGroupName;

		IF @ApplicationGroupId IS NOT NULL
			INSERT INTO IisSiteApplicationGroupAssociations (
				IisSiteId
				,IisApplicationGroupId
				)
			SELECT a.UniqueId AS IisSiteId
				,@ApplicationGroupId
			FROM @IisSiteIds a;
		ELSE
		BEGIN
			SELECT @ErrorMessage = CONCAT (
					'Application group '
					,@ApplicationGroupName
					,' does not exist'
					);

			THROW 51000
				,@ErrorMessage
				,1;
		END;
	END;

	COMMIT TRANSACTION CreateSites;

	RETURN 0
END;
