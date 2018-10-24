using Microsoft.EntityFrameworkCore;

namespace PlatformCI.InfrastructureService.Entities.EfContextBuilder
{
    public class EfContextBuilder
    {
        public EfContextBuilder(DbContextOptions<BuildSystemContext> dbContextOptions)
        {
            DbContextOptions = dbContextOptions;
        }

        public DbContextOptions<BuildSystemContext> DbContextOptions { get; }

        public BuildSystemContext GetNewContext()
        {
            return new BuildSystemContext(DbContextOptions);
        }
    }
}