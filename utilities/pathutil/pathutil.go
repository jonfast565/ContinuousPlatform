// utility for dealing with paths
package pathutil

import (
	"github.com/ahmetb/go-linq"
	"strings"
)

// path constants
const goBack string = ".."
const stay string = "."

var pathSplitterChar = '/'

// Normalizes a path into Unix path format (instead of Windows)
// The Unix format is easier to work with
func NormalizePath(str string) string {
	result := strings.Replace(str, "\\", "/", -1)
	return result
}

// A path action is an activity to be done on a path.
// A child is traversed with a 'name'
// The current folder is traversed with a '.'
// The parent folder is traversed with a '..'
type PathAction struct {
	Name string
}

// List of actions to take, coming from a URI
type PathActionList []PathAction

// A parser of a path that will produce a specific action series
type PathParser struct {
	ActionSeries *[]PathAction
}

// Gets the last value in a series of values representing a path
func GetLastPathComponent(path string) string {
	parser := NewPathParserFromString(path)
	lastItem := parser.GetLastItem()
	return lastItem
}

// Takes a path string and creates a new path parser with action series children
func NewPathParserFromString(path string) *PathParser {
	parser := new(PathParser)
	parser.SetActionSeries(path)
	return parser
}

// Adds an action that traverses into a child folder
func (parser *PathParser) AddGoAheadAction(pathFragment string) {
	*parser.ActionSeries = append(*parser.ActionSeries, PathAction{Name: pathFragment})
}

// Sets the action series (really does parsing)
func (parser *PathParser) SetActionSeries(path string) {
	items := make([]PathAction, 0)
	if path == "" {
		parser.ActionSeries = &items
		return
	}
	normalizedPath := NormalizePath(path)
	splitFn := func(c rune) bool {
		return c == pathSplitterChar
	}
	splitPath := strings.FieldsFunc(normalizedPath, splitFn)
	for _, pathPart := range splitPath {
		switch pathPart {
		case goBack:
			items = append(items, PathAction{Name: ".."})
			break
		case stay:
			items = append(items, PathAction{Name: "."})
			break
		default:
			items = append(items, PathAction{Name: pathPart})
		}
	}
	parser.ActionSeries = &items
}

// Gets the last item in a path
func (parser *PathParser) GetLastItem() string {
	if len(*parser.ActionSeries) == 0 {
		panic("No values in this path")
	}
	result := linq.From(*parser.ActionSeries).SelectT(
		func(iterator PathAction) string {
			return iterator.Name
		}).Last()

	if str, ok := result.(string); ok {
		return str
	} else {
		panic("Path value not a string")
	}
}

// Gets all action series items in a path as a string
func (parser *PathParser) GetAllItems() []string {
	results := linq.From(*parser.ActionSeries).SelectT(
		func(iterator PathAction) string {
			return iterator.Name
		})

	var result []string
	results.ToSlice(&result)
	return result
}

// Gets all action series items before the final item in the action series
func (parser *PathParser) GetPreviousItems() []string {
	results := linq.From(*parser.ActionSeries).SelectT(
		func(iterator PathAction) string {
			return iterator.Name
		})

	var result []string
	results.ToSlice(&result)
	if len(result)-1 == 0 {
		return []string{}
	} else {
		sliceLast := result[:len(result)-1]
		return sliceLast
	}
}

// Gets a path string from an action series, with a customizable start delimiter.
func (parser *PathParser) GetPathString(includeStartDelimiter bool) string {
	var result = ""
	if includeStartDelimiter {
		result += "./"
	}
	if !linq.From(*parser.ActionSeries).Any() {
		return result
	}
	var seriesList = *parser.ActionSeries
	for counter, action := range seriesList {
		result += action.Name
		if counter != len(seriesList)-1 {
			result += "/"
		}
	}
	return result
}

// An object for zipping two paths together
type PathActionZipped struct {
	item1 PathAction
	item2 PathAction
}

// Zip list of path actions
type PathActionZipList []PathActionZipped

// Determines if a partial match of two paths exist (for instance, parent child relationship)
func (pazl PathActionZipList) PartialMatch() bool {
	for _, zipAction := range pazl {
		if zipAction.item1 != zipAction.item2 {
			return false
		}
	}
	return true
}

// Zips two path parsers together in one list
func (parser *PathParser) ZipPathParsers(parser2 *PathParser) PathActionZipList {
	zippedSeries := make(PathActionZipList, 0)
	pLength := len(*parser.ActionSeries)
	for counter, item := range *parser.ActionSeries {
		if counter > pLength-1 {
			break
		}
		item2 := (*parser2.ActionSeries)[counter]
		zippedSeries = append(zippedSeries, PathActionZipped{
			item1: item,
			item2: item2,
		})
	}
	return zippedSeries
}
