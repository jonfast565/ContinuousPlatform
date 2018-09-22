package utilities

import (
	"github.com/ahmetb/go-linq"
	"strings"
)

// path constants
const goBack string = ".."
const stay string = "."
const server string = "$"
const shebangHalf string = "!"

var pathSplitterChar rune = '/'

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
	Series *PathActionSeries
}

type PathActionSeries struct {
	ActionItems *[]Namer
}

func GetLastPathComponent(path string) string {
	parser := new(PathParser)
	parser.SetActionSeries(path)
	lastItem := parser.GetLastItem()
	return lastItem
}

func (actionSeries *PathActionSeries) New(actionItems *[]Namer) {
	actionSeries.ActionItems = actionItems
}

func (parser *PathParser) SetActionSeries(path string) {
	items := make([]Namer, 0)
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
	result := PathActionSeries{ActionItems: &items}
	parser.Series = &result
}

func (parser *PathParser) GetLastItem() string {
	result := linq.From(parser.Series.ActionItems).SelectT(
		func(iterator Namer) string {
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
	if !linq.From(parser.Series).Any() {
		return result
	}
	var seriesList = *parser.Series.ActionItems
	for counter, action := range seriesList {
		result += action.GetName()
		if counter != len(seriesList)-1 {
			result += "/"
		}
	}
	return result
}

type PathActionZipped struct {
	item1 Namer
	item2 Namer
}

func NullZipSeries(series1 PathActionSeries, series2 PathActionSeries, strictLength bool) []PathActionZipped {
	zippedSeries := make([]PathActionZipped, 0)
	items1Length := len(*series1.ActionItems)
	for counter, item1 := range *series1.ActionItems {
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
				item2: (*series2.ActionItems)[counter],
			})
		}
	}
	return zippedSeries
}

func (parser *PathParser) ContainsOrEquals(series2 PathActionSeries, strictLength bool) {

}
