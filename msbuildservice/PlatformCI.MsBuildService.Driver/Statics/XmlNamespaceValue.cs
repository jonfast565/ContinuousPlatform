using System.Xml;

namespace BuildSystem.Lib.MicrosoftBuildProvider.Statics
{
    public class XmlNamespaceValue
    {
        public XmlNamespaceManager NsManager { get; set; }
        public string NamespacePrefix { get; set; }
        public string NamespaceName { get; set; }
    }
}