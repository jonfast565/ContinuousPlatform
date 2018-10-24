namespace PlatformCI.InfrastructureService.Models.Interfaces
{
    public interface IRepositorySolutionProjectKey
    {
        string RepositoryName { get; set; }
        string SolutionName { get; set; }
        string ProjectName { get; set; }
    }
}