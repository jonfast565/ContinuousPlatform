using System.Collections.Generic;
using PlatformCI.InfrastructureService.Models.Implementation;

namespace PlatformCI.InfrastructureService.Models.Comparers
{
    public class IisApplicationPoolEqualityComparer : IEqualityComparer<IisApplicationPool>
    {
        public bool Equals(IisApplicationPool x, IisApplicationPool y)
        {
            return x.AppPoolGuid == y.AppPoolGuid;
        }

        public int GetHashCode(IisApplicationPool obj)
        {
            return 0;
        }
    }
}