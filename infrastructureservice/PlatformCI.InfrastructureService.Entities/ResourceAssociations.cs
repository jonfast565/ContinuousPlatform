using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class ResourceAssociations
    {
        public Guid ResourceAssociationId { get; set; }
        public Guid DeliverableId { get; set; }
        public Guid? IisSiteId { get; set; }
        public Guid? IisApplicationId { get; set; }
        public Guid? WindowsServiceId { get; set; }
        public Guid? ScheduledTaskId { get; set; }
        public Guid? DeploymentLocationId { get; set; }
        public Guid ResourceAmalgamationId { get; set; }
        public bool? Enabled { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public Deliverables Deliverable { get; set; }
        public DeploymentLocations DeploymentLocation { get; set; }
        public IisApplications IisApplication { get; set; }
        public IisSites IisSite { get; set; }
        public ResourceAmalgamations ResourceAmalgamation { get; set; }
        public ScheduledTasks ScheduledTask { get; set; }
        public WindowsServices WindowsService { get; set; }
    }
}