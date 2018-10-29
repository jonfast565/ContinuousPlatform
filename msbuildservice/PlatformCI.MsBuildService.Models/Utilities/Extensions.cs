using System;
using System.Collections.Generic;
using System.Text;

namespace PlatformCI.MsBuildService.Models.Utilities
{
    public static class Extensions
    {
        public static string NormalizePath(this string pathString)
        {
            return pathString.Replace('\\', '/');
        }
    }
}
