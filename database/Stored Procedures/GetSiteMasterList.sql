CREATE PROCEDURE GetSiteMasterList
AS
BEGIN
	SELECT 
		e.EnvironmentName,
		--isg.GroupName as SiteGroup,
		s.SiteName as SiteName,
		s.PhysicalPath as SitePhysicalPath,
        iaps.PoolName as SiteApplicationPool,
		--iag.GroupName as ApplicationGroupName,
		ia.ApplicationName as ApplicationName,
		REPLACE(ia.PhysicalPath, '{{SiteName}}', s.SiteName) as ApplicationPhysicalPath,
		iap.PoolName as ApplicationAppPoolName
	FROM Environments e
	inner join IisSiteGroups isg on isg.EnvironmentId = e.EnvironmentId
	inner join IisSiteGroupAssociations isga on isga.IisSiteGroupId = isg.IisSiteGroupId
	inner join IisSites s on s.IisSiteId = isga.IisSiteId
    inner join IisApplicationPools iaps on s.IisSiteApplicationPoolId = iaps.IisApplicationPoolId
	left outer join IisSiteApplicationGroupAssociations isaga on isaga.IisSiteId = isga.IisSiteId
	left outer join IisApplicationGroups iag on iag.IisApplicationGroupId = isaga.IisApplicationGroupId
	left outer join IisApplicationGroupAssociations iaga on iaga.IisApplicationGroupId = iag.IisApplicationGroupId
	left outer join IisApplications ia on ia.IisApplicationId = iaga.IisApplicationId
	left outer join IisApplicationPools iap on ia.IisApplicationPoolId = iap.IisApplicationPoolId;
    RETURN 0;
END
GO