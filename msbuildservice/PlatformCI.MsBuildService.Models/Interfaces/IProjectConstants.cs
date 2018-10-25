using System.Text.RegularExpressions;

namespace BuildSystem.Lib.Models.Deliverable.Interfaces
{
    public interface IProjectConstants
    {
        Regex[] AllValidExtensions { get; }
        Regex[] ValidProjectExtensions { get; }
        Regex[] ValidSolutionFileExtensions { get; }
        Regex[] ValidPublishProfileExtensions { get; }
        string PublishProfileProjectSubfolderPath { get; }
    }
}