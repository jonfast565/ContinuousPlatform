using System.Collections.Generic;

namespace BuildSystem.Lib.PathParser.Interfaces
{
    public interface IPathActionSeries
    {
        ICollection<ISourceControlPathAction> Actions { get; }
        string GetPathString(bool includeStartDelimiter = true);
        bool ContainsOrEquals(IPathActionSeries series, bool strictEquals);
        IPathActionSeries MakeActionSeriesRelativeTo(IPathActionSeries series);
        string GetLastItem();
        IPathActionSeries RemoveLastNActions(int numItems);
    }
}