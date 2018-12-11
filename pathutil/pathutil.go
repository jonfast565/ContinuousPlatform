package pathutil

import (
	"github.com/ahmetb/go-linq"
	"strings"
)

// path constants
const goBack string = ".."
const stay string = "."

var pathSplitterChar = '/'

func NormalizePath(str string) string {
	result := strings.Replace(str, "\\", "/", -1)
	return result
}

type PathAction struct {
	Name string
}

type PathActionList []PathAction

type PathParser struct {
	ActionSeries *[]PathAction
}

func GetLastPathComponent(path string) string {
	parser := NewPathParserFromString(path)
	lastItem := parser.GetLastItem()
	return lastItem
}

func NewPathParserFromString(path string) *PathParser {
	parser := new(PathParser)
	parser.SetActionSeries(path)
	return parser
}

func (parser *PathParser) AddGoAheadAction(pathFragment string) {
	*parser.ActionSeries = append(*parser.ActionSeries, PathAction{Name: pathFragment})
}

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

func (parser *PathParser) GetAllItems() []string {
	results := linq.From(*parser.ActionSeries).SelectT(
		func(iterator PathAction) string {
			return iterator.Name
		})

	var result []string
	results.ToSlice(&result)
	return result
}

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

type PathActionZipped struct {
	item1 PathAction
	item2 PathAction
}

type PathActionZipList []PathActionZipped

func (pazl PathActionZipList) PartialMatch() bool {
	for _, zipAction := range pazl {
		if zipAction.item1 != zipAction.item2 {
			return false
		}
	}
	return true
}

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
