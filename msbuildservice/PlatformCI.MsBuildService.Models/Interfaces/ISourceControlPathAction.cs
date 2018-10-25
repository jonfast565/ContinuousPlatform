using BuildSystem.Lib.PathParser.Enums;

namespace BuildSystem.Lib.PathParser.Interfaces
{
    public interface ISourceControlPathAction
    {
        string NextDirectory { get; }
        PathActionType Action { get; }
    }
}