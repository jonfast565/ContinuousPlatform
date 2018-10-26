using System.Collections.Generic;
using PlatformCI.MsBuildService.Models.Implementation;

namespace PlatformCI.MsBuildService.Models.Interfaces
{
    public interface IFilesystemProvider
    {
        string TempFolderPath { get; set; }
        bool DirectoryEmpty(string absolutePath);
        bool DirectoryExists(string absolutePath);
        void CreateDirectory(string absolutePath, bool createIfNotExists = true);
        void DeleteDirectory(string absolutePath);
        void CopyFile(string sourceAbsolutePath, string destinationAbsolutePath);
        void DeleteFile(string absolutePath);
        bool FileExists(string absolutePath);
        string GetTemporarySystemPath();
        void CleanTemporarySystemPath();
        string ReadFile(string absolutePath);
        void WriteFile(string absolutePath, string contents);
        void WriteFile(string absolutePath, byte[] data);
        IList<FilesystemMetadata> GetContents(string absolutePath, bool recursiveSearch, bool foldersOnly);
        string GetParentPath(string absolutePath);
        string GetSecondParentPath(string absolutePath);
        void UnsetReadOnlyAttributesForDirectory(string absolutePath);
        void UnsetReadOnlyAttributesForFile(string absolutePath);
        void CopyDirectory(string source, string destination);
        byte[] ReadFileBytes(string v);
        void SetCurrentDirectory(string directoryName);
    }
}