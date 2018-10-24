using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class IisSiteBindings
    {
        public Guid IisSiteBindingId { get; set; }
        public Guid IisSiteId { get; set; }
        public string BindingName { get; set; }
        public int Port { get; set; }
        public string CertificateThumbprint { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public IisSites IisSite { get; set; }
    }
}