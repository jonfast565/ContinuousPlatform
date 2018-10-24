using System.Collections.Generic;
using PlatformCI.InfrastructureService.Models.Implementation;

namespace PlatformCI.InfrastructureService.Models.Comparers
{
    public class IisSiteEqualityComparer : IEqualityComparer<IisSite>
    {
        public bool Equals(IisSite x, IisSite y)
        {
            return x.SiteGuid == y.SiteGuid;
        }

        public int GetHashCode(IisSite obj)
        {
            return 0;
        }
    }
}