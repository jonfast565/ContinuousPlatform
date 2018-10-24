using System.Collections.Generic;
using PlatformCI.InfrastructureService.Models.Implementation;

namespace PlatformCI.InfrastructureService.Models.Comparers
{
    internal class WindowsServiceEqualityComparer : IEqualityComparer<WindowsService>
    {
        public bool Equals(WindowsService x, WindowsService y)
        {
            return x.ServiceGuid == y.ServiceGuid;
        }

        public int GetHashCode(WindowsService obj)
        {
            return 0;
        }
    }
}