using System;
using System.Collections.Generic;
using System.Linq;
using BuildSystem.Lib.PathParser.Enums;
using BuildSystem.Lib.PathParser.Interfaces;

namespace BuildSystem.Lib.PathParser.Implementation
{
    public class PathActionSeries
    {
        public PathActionSeries(ICollection<ISourceControlPathAction> actions)
        {
            Actions = actions;
        }

        public ICollection<ISourceControlPathAction> Actions { get; }
    }
}