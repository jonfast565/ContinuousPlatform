using System;
using System.Collections.Generic;
using System.Linq;
using PlatformCI.InfrastructureService.Models.Comparers;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class IisSiteApplicationMemberPool
    {
        public IisSite ParentSite { get; set; }
        public IList<IisApplication> ChildApplications { get; set; }

        public static IList<IisSiteApplicationMemberPool> InitPools(IList<IisSite> sites,
            IList<IisApplication> applications)
        {
            return (from site in sites
                let siteAppList = applications.Where(x => x.Sites.Contains(site, new IisSiteEqualityComparer()))
                    .ToList()
                select new IisSiteApplicationMemberPool
                {
                    ParentSite = site,
                    ChildApplications = siteAppList.Select(app => new IisApplication
                    {
                        ApplicationGuid = app.ApplicationGuid,
                        ApplicationName = app.ApplicationName,
                        AppPool = app.AppPool,
                        PhysicalPath = app.PhysicalPath.Replace("{{SiteName}}", site.SiteName),
                        Sites = app.Sites
                    }).ToList()
                }).ToList();
        }
    }
}