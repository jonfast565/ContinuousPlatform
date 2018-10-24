using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class EnvironmentInfrastructure
    {
        public string Environment { get; set; }
        public IList<IisApplicationPool> IisApplicationPools { get; set; }
        public IList<ScheduledTask> ScheduledTasks { get; set; }
        public IList<WindowsService> WindowsServices { get; set; }
        public IList<IisSiteApplicationMemberPool> SiteApplicationMembers { get; set; }
        public IList<ServerNameType> Servers { get; set; }
    }
}