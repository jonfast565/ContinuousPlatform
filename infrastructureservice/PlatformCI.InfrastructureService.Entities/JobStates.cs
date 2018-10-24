using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class JobStates
    {
        public JobStates()
        {
            Jobs = new HashSet<Jobs>();
        }

        public Guid JobStateId { get; set; }
        public string StateName { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<Jobs> Jobs { get; set; }
    }
}