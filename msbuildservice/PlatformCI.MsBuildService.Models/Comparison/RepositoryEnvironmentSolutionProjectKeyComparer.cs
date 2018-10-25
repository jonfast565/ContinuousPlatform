using System.Collections.Generic;
using BuildSystem.Lib.Models.Deliverable.Implementation;

namespace BuildSystem.Lib.Models.Deliverable.Comparison
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