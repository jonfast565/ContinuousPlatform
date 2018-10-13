package pathutil

import (
	"../interfaces"
	"github.com/ahmetb/go-linq"
	"strings"
)

// path constants
const goBack string = ".."
const stay string = "."
const server string = "$"
const shebangHalf string = "!"

var pathSplitterChar = '/'

func NormalizePath(str string) string {
	result := strings.Replace(str, "\\", "/", -1)
	return result
}

type PathActionGoAhead struct {
	Name string
}

func (pa PathActionGoAhead) GetName() string {
	return pa.Name
}

type PathActionGoBack struct {
}

func (pa PathActionGoBack) GetName() string {
	return ".."
}

type PathActionStay struct {
}

func (pa PathActionStay) GetName() string {
	return "."
}

type PathParser struct {
	ActionSeries *[]interfaces.Namer
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
	*parser.ActionSeries = append(*parser.ActionSeries, PathActionGoAhead{Name: pathFragment})
}

func (parser *PathParser) SetActionSeries(path string) {
	items := make([]interfaces.Namer, 0)
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
			items = append(items, PathActionGoBack{})
			break
		case stay:
		case server:
		case shebangHalf:
			items = append(items, PathActionStay{})
			break
		default:
			items = append(items, PathActionGoAhead{Name: pathPart})
		}
	}
	parser.ActionSeries = &items
}

func (parser *PathParser) GetLastItem() string {
	result := linq.From(*parser.ActionSeries).SelectT(
		func(iterator interfaces.Namer) string {
			return iterator.GetName()
		}).Last()

	if str, ok := result.(string); ok {
		return str
	} else {
		panic("Not a string")
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
		result += action.GetName()
		if counter != len(seriesList)-1 {
			result += "/"
		}
	}
	return result
}

type PathActionZipped struct {
	item1 interfaces.Namer
	item2 interfaces.Namer
}

func (parser *PathParser) NullZipSeries(parser2 *PathParser, strictLength bool) []PathActionZipped {
	zippedSeries := make([]PathActionZipped, 0)
	items1Length := len(*parser.ActionSeries)
	for counter, item1 := range *parser.ActionSeries {
		if counter >= items1Length-1 && strictLength {
			break
		} else if counter >= items1Length-1 && !strictLength {
			zippedSeries = append(zippedSeries, PathActionZipped{
				item1: item1,
				item2: nil,
			})
		} else {
			zippedSeries = append(zippedSeries, PathActionZipped{
				item1: item1,
				item2: (*parser2.ActionSeries)[counter],
			})
		}
	}
	return zippedSeries
}

func (parser *PathParser) RemoveLastNActions(nActions int) {
	if len(*parser.ActionSeries) <= 0 {
		return
	}
	for i := 0; i < nActions; i++ {
		if len(*parser.ActionSeries) > 0 {
			*parser.ActionSeries = (*parser.ActionSeries)[:len(*parser.ActionSeries)-1]
		}
	}
}
