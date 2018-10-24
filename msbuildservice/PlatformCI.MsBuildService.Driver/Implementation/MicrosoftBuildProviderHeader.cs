using BuildSystem.Lib.Interfaces.Generic.Interfaces;
using BuildSystem.Lib.Utils.Statics;

namespace BuildSystem.Lib.MicrosoftBuildProvider.Implementation
{
    internal class MicrosoftBuildProviderHeader : IHeader
    {
        public void PrintHeader()
        {
            ApplicationHeader.PrintApplicationHeader("Microsoft Build Provider");
        }
    }
}