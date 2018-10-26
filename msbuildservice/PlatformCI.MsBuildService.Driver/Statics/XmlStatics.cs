using System.Collections.Generic;
using System.Linq;
using System.Xml;

namespace PlatformCI.MsBuildService.Driver.Statics
{
    public static class XmlStatics
    {
        public static void RemoveComments(this XmlDocument document)
        {
            foreach (var node in document.SelectNodes("//comment()").Cast<XmlNode>().ToList())
                node.ParentNode.RemoveChild(node);
        }

        public static XmlNamespaceValue CreateLegacyMsBuildNamespace(this XmlDocument document)
        {
            var namespaceName = "ns";
            var namespacePrefix = string.Empty;
            XmlNamespaceManager nameSpaceManager = null;
            document.RemoveComments();

            var firstChildNode = document.ChildNodes[1];
            if (firstChildNode?.Attributes != null)
            {
                var xmlns = firstChildNode.Attributes["xmlns"];
                if (xmlns != null)
                {
                    nameSpaceManager = new XmlNamespaceManager(document.NameTable);
                    nameSpaceManager.AddNamespace(namespaceName, xmlns.Value);
                    namespacePrefix = namespaceName + ":";
                }
            }

            return new XmlNamespaceValue
            {
                NsManager = nameSpaceManager,
                NamespacePrefix = namespacePrefix,
                NamespaceName = namespaceName
            };
        }

        public static string GetSingleNodeInnerText(
            this XmlDocument document,
            string xpathArgument,
            XmlNamespaceValue nsValue)
        {
            if (nsValue != null)
            {
                var node = document.SelectSingleNode(xpathArgument, nsValue.NsManager);
                return node?.InnerText;
            }
            else
            {
                var node = document.SelectSingleNode(xpathArgument);
                return node?.InnerText;
            }
        }

        public static IEnumerable<string> GetNodesPropertyValue(
            this XmlDocument document,
            string xpathArgument,
            string propertyValue,
            XmlNamespaceValue nsValue)
        {
            var list = document.SelectNodes(xpathArgument, nsValue.NsManager);
            return list.Cast<XmlNode>().Select(x => x.Attributes[propertyValue].Value);
        }
    }
}