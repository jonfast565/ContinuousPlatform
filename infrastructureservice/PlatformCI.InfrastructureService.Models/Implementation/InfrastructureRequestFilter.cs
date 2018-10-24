using System;
using System.Collections.Generic;
using System.Text;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    public class InfrastructureRequestFilter
    {
        public string RepositoryName { get; set; }
        public string SolutionName { get; set; }
        public string ProjectName { get; set; }
        public string OptionalEnvironmentName { get; set; }
    }
}
