using System.Collections.Generic;
using PlatformCI.InfrastructureService.Models.Implementation;

namespace PlatformCI.InfrastructureService.Models.Comparers
{
    public class ScheduledTaskEqualityComparer : IEqualityComparer<ScheduledTask>
    {
        public bool Equals(ScheduledTask x, ScheduledTask y)
        {
            return x.TaskGuid == y.TaskGuid;
        }

        public int GetHashCode(ScheduledTask obj)
        {
            return 0;
        }
    }
}