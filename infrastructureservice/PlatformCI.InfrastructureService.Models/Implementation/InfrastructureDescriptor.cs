using System;
using System.Collections.Generic;
using System.Linq;
using PlatformCI.InfrastructureService.Models.Interfaces;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class InfrastructureDescriptor : IRepositorySolutionProjectKey
    {
        public IList<IisSite> IisSites { get; set; }
        public IList<IisApplication> IisApplications { get; set; }
        public IList<IisApplicationPool> IisApplicationPools { get; set; }
        public IList<WindowsService> WindowsServices { get; set; }
        public IList<ScheduledTask> ScheduledTasks { get; set; }
        public IList<string> ApplicableEnvironments { get; set; }
        public string RepositoryName { get; set; }
        public string SolutionName { get; set; }
        public string ProjectName { get; set; }

        public bool AnyPhysicalInfrastructure()
        {
            return IisSites.Any()
                   || IisApplications.Any()
                   || WindowsServices.Any()
                   || ScheduledTasks.Any();
        }
    }
}