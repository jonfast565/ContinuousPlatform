using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class Logs
    {
        public Guid LogId { get; set; }
        public DateTime Date { get; set; }
        public string MachineName { get; set; }
        public string ApplicationName { get; set; }
        public string LogLevel { get; set; }
        public string Message { get; set; }
    }
}