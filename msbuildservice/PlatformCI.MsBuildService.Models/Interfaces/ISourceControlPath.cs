using System.Collections.Generic;

namespace BuildSystem.Lib.PathParser.Interfaces
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