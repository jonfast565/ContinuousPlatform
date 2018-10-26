using System.Collections.Generic;
using PlatformCI.MsBuildService.Models.Implementation;

namespace PlatformCI.MsBuildService.Models.Comparison
{
    public class
        RepositoryEnvironmentSolutionProjectKeyComparer : IEqualityComparer<RepositoryBranchSolutionProjectKey>
    {
        public bool Equals(RepositoryBranchSolutionProjectKey x, RepositoryBranchSolutionProjectKey y)
        {
            return x.Repository == y.Repository && x.Branch == y.Branch && x.Solution == y.Solution &&
                   x.Project == y.Project;
        }

        public int GetHashCode(RepositoryBranchSolutionProjectKey obj)
        {
            return obj.GetHashCode();
        }
    }
}