using System.Collections.Generic;
using System.Linq;
using PlatformCI.MsBuildService.Models.Abstractions;
using PlatformCI.MsBuildService.Models.Interfaces;

namespace PlatformCI.MsBuildService.Models.Implementation
{
    public class PathActionSeries
    {
        public PathActionSeries(ICollection<SourceControlPathAction> actions)
        {
            Actions = actions;
        }

        public ICollection<SourceControlPathAction> Actions { get; }

        public string GetLastItem()
        {
            return Actions.LastOrDefault()?.NextDirectory;
        }
    }
}