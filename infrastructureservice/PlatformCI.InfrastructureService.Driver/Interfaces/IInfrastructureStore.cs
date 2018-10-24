using PlatformCI.InfrastructureService.Models.Implementation;

namespace PlatformCI.InfrastructureService.Driver.Interfaces
{
    public interface IInfrastructureStore
    {
        InfrastructureResult GetInfrastructureMetadata(InfrastructureRequestFilter requestFilter);
    }
}