using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class WindowsServices
    {
        public WindowsServices()
        {
            ResourceAssociations = new HashSet<ResourceAssociations>();
            WindowsServiceGroupAssociations = new HashSet<WindowsServiceGroupAssociations>();
        }

        public Guid WindowsServiceId { get; set; }
        public string ServiceName { get; set; }
        public string BinaryPath { get; set; }
        public string BinaryExecutableName { get; set; }
        public string BinaryExecutableArguments { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<ResourceAssociations> ResourceAssociations { get; set; }
        public ICollection<WindowsServiceGroupAssociations> WindowsServiceGroupAssociations { get; set; }
    }
}