using System.Collections.Generic;
using PlatformCI.InfrastructureService.Models.Implementation;

namespace PlatformCI.InfrastructureService.Models.Comparers
{
    public class IisApplicationEqualityComparer : IEqualityComparer<IisApplication>
    {
        public bool Equals(IisApplication x, IisApplication y)
        {
            return x.ApplicationGuid == y.ApplicationGuid;
        }

        public int GetHashCode(IisApplication obj)
        {
            return 0;
        }
    }
}