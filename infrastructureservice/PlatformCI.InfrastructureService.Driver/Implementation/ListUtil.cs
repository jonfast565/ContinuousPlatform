using System.Collections.Generic;
using System.Linq;

namespace PlatformCI.InfrastructureService.Driver.Implementation
{
    public static class ListUtil
    {
        public static List<T> ConcatLists<T>(params IEnumerable<T>[] lists)
        {
            return lists.Aggregate(new List<T>(), (current, list) => current.Concat(list).ToList()).ToList();
        }
    }
}
