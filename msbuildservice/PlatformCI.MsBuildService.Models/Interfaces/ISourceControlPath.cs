using System.Collections.Generic;

namespace PlatformCI.MsBuildService.Models.Interfaces
{
    public interface ISourceControlPath
    {
        ICollection<ISourceControlPathAction> GetActionList();
        void ReplenishActionQueue();
        ISourceControlPathAction Dequeue();
        string PathString();
        bool AnyActions();
        int NumberOfRemainingActions();
    }
}