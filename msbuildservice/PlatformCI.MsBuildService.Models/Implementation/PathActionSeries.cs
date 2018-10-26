using System.Collections.Generic;
using PlatformCI.MsBuildService.Models.Interfaces;

namespace PlatformCI.MsBuildService.Models.Implementation
{
    public class PathActionSeries
    {
        public PathActionSeries(ICollection<ISourceControlPathAction> actions)
        {
            Actions = actions;
        }

        public ICollection<ISourceControlPathAction> Actions { get; }
    }
}