using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using PlatformCI.MsBuildService.Models.Enums;
using PlatformCI.MsBuildService.Models.Interfaces;

namespace PlatformCI.MsBuildService.Models.Implementation
{
    public class BasicFilesystemProvider : IFilesystemProvider
    {
        public BasicFilesystemProvider(string tempFolderPath = @"C:\Temp\")
        {
            TempFolderPath = tempFolderPath;
            CreateDirectory(TempFolderPath);
        }

        public string TempFolderPath { get; set; }

        public bool DirectoryExists(string absolutePath)
        {
            return Directory.Exists(absolutePath);
        }

        public void CreateDirectory(string absolutePath, bool createIfNotExists = true)
        {
            if (createIfNotExists)
            {
                Directory.CreateDirectory(absolutePath);
                return;
            }

            if (!DirectoryExists(absolutePath))
                Directory.CreateDirectory(absolutePath);
            else
                throw new Exception($"File {absolutePath} does not exist");
        }

        public void DeleteDirectory(string absolutePath)
        {
            if (DirectoryExists(absolutePath))
                try
                {
                    Directory.Delete(absolutePath, true);
                }
                catch (IOException)
                {
                    // all I'm saying is, WTF Microsoft?!
                    // why does this work? I have no idea why this works...
                    // Thread.Sleep(100);
                    Directory.Delete(absolutePath, true);
                }
            else
                throw new Exception($"File {absolutePath} does not exist");
        }

        public void CopyFile(string sourceAbsolutePath, string destinationAbsolutePath)
        {
            if (FileExists(sourceAbsolutePath))
                File.Copy(sourceAbsolutePath, destinationAbsolutePath);
            else
                throw new Exception($"File {sourceAbsolutePath} does not exist");
        }

        public void DeleteFile(string absolutePath)
        {
            File.Delete(absolutePath);
        }

        public bool FileExists(string absolutePath)
        {
            return File.Exists(absolutePath);
        }

        public string GetTemporarySystemPath()
        {
            return TempFolderPath + Guid.NewGuid() + ".tmp";
        }

        public void CleanTemporarySystemPath()
        {
            var di = new DirectoryInfo(TempFolderPath);

            if (!di.Exists) return;

            foreach (var file in di.EnumerateFiles())
                if (FileExists(file.FullName))
                    DeleteFile(file.FullName);

            foreach (var folder in di.EnumerateDirectories())
                if (DirectoryExists(folder.FullName))
                    DeleteDirectory(folder.FullName);
        }

        public string ReadFile(string absolutePath)
        {
            if (FileExists(absolutePath))
                return File.ReadAllText(absolutePath);
            throw new Exception($"File {absolutePath} does not exist.");
        }

        public void WriteFile(string absolutePath, string contents)
        {
            File.WriteAllText(absolutePath, contents);
        }

        public void WriteFile(string absolutePath, byte[] bytes)
        {
            File.WriteAllBytes(absolutePath, bytes);
        }

        public IList<FilesystemMetadata> GetContents(string absolutePath, bool recursiveSearch, bool foldersOnly)
        {
            var searchWildcard = "*";
            var searchOption = recursiveSearch ? SearchOption.AllDirectories : SearchOption.TopDirectoryOnly;
            var fileFolderList = new List<FilesystemMetadata>();

            var dirInfo = new DirectoryInfo(absolutePath);
            var dirs = dirInfo.EnumerateDirectories(searchWildcard, searchOption);

            fileFolderList.AddRange(dirs.Select(x =>
                new FilesystemMetadata
                {
                    Path = x.FullName,
                    Type = FilesystemObjectType.Folder
                }));

            if (!foldersOnly)
            {
                var files = dirInfo.EnumerateFiles(searchWildcard, searchOption);
                fileFolderList.AddRange(files.Select(x =>
                    new FilesystemMetadata
                    {
                        Path = x.FullName,
                        Type = FilesystemObjectType.File
                    }));
            }

            return fileFolderList;
        }

        public string GetParentPath(string absolutePath)
        {
            var directoryInfo = new DirectoryInfo(absolutePath).Parent;
            return directoryInfo?.FullName;
        }

        public string GetSecondParentPath(string absolutePath)
        {
            return GetParentPath(GetParentPath(absolutePath));
        }

        public bool DirectoryEmpty(string absolutePath)
        {
            return !new DirectoryInfo(absolutePath).GetFileSystemInfos().Any();
        }

        public void UnsetReadOnlyAttributesForDirectory(string absolutePath)
        {
            var di = new DirectoryInfo(absolutePath);
            di.Attributes &= ~FileAttributes.ReadOnly;
        }

        public void UnsetReadOnlyAttributesForFile(string absolutePath)
        {
            var fi = new FileInfo(absolutePath);
            fi.Attributes &= ~FileAttributes.ReadOnly;
        }

        public void CopyDirectory(string sourceDirectory, string targetDirectory)
        {
            void All(DirectoryInfo source, DirectoryInfo target)
            {
                Directory.CreateDirectory(target.FullName);

                // Copy each file into the new directory.
                foreach (var fi in source.GetFiles())
                    fi.CopyTo(Path.Combine(target.FullName, fi.Name), true);

                // Copy each subdirectory using recursion.
                foreach (var diSourceSubDir in source.GetDirectories())
                {
                    var nextTargetSubDir = target.CreateSubdirectory(diSourceSubDir.Name);
                    All(diSourceSubDir, nextTargetSubDir);
                }
            }

            var diSource = new DirectoryInfo(sourceDirectory);
            var diTarget = new DirectoryInfo(targetDirectory);

            All(diSource, diTarget);
        }

        public byte[] ReadFileBytes(string absolutePath)
        {
            if (FileExists(absolutePath))
                return File.ReadAllBytes(absolutePath);
            throw new Exception($"File {absolutePath} does not exist.");
        }

        public void SetCurrentDirectory(string directoryName)
        {
            Directory.SetCurrentDirectory(directoryName);
        }
    }
}