using System;

namespace BuildSystem.Lib.Models.Deliverable.Implementation
{
    [Serializable]
    public class RepositoryBranchSolutionProjectKey
    {
        public string Repository { get; set; }
        public string Branch { get; set; }
        public string Solution { get; set; }
        public string Project { get; set; }
    }
}