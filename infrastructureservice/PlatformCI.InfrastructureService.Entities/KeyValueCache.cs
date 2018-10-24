using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class KeyValueCache
    {
        public Guid KeyValueCacheId { get; set; }
        public string Key { get; set; }
        public byte[] Value { get; set; }
        public string ValueType { get; set; }
        public string MachineName { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }
    }
}