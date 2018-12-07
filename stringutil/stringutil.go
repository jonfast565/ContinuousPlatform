package stringutil

import (
	"regexp"
	"strings"
)

func ConcatMultipleWithSeparator(separator string, inputs ...string) string {
	var result string
	for i, value := range inputs {
		result += value
		if i != len(inputs)-1 {
			result += separator
		}
	}
	return result
}

func ConcatMultiple(inputs ...string) string {
	return ConcatMultipleWithSeparator(" ", inputs...)
}

func ConcatDelimitMultiple(separator string, leftDelimiter string, rightDelimiter string, inputs []string) string {
	var result string
	for i, value := range inputs {
		result += leftDelimiter + value + rightDelimiter
		if i != len(inputs)-1 {
			result += separator
		}
	}
	return result
}

func CompileStringsAsRegexes(regexStrings []string) ([]regexp.Regexp, error) {
	results := make([]regexp.Regexp, 0)
	for _, regexString := range regexStrings {
		regexValue, err := regexp.Compile(regexString)
		if err != nil {
			return nil, err
		}
		results = append(results, *regexValue)
	}
	return results, nil
}

func StringMatchesOneOfRegStr(value string, comparators []string) (bool, error) {
	compiledRegexes, err := CompileStringsAsRegexes(comparators)
	if err != nil {
		return false, err
	}
	match := StringMatchesOneOf(value, compiledRegexes)
	return match, nil
}

func StringMatchesOneOf(value string, comparators []regexp.Regexp) bool {
	for _, comparator := range comparators {
		if comparator.Match([]byte(value)) {
			return true
		}
	}
	return false
}

func PartialMessage(value string) string {
	maxMessageLength := 40
	if len(value) > maxMessageLength {
		maxMessageLength = len(value)
	}
	return value[0 : maxMessageLength-1]
}

func StringArrayContains(strArray []string, value string) bool {
	for _, str := range strArray {
		if str == value {
			return true
		}
	}
	return false
}

func StringArrayCompare(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, el1 := range arr1 {
		if arr2[i] != el1 {
			return false
		}
	}
	return true
}

func StringArrayCompareNumeric(arr1, arr2 []string) int {
	if len(arr1) != len(arr2) {
		if len(arr1) < len(arr2) {
			return -1
		} else {
			return 1
		}
	}

	for i, el1 := range arr1 {
		if comp := strings.Compare(arr2[i], el1); comp != 0 {
			return comp
		}
	}
	return 0
}

func StringArrayContainsArray(arr1, arr2 []string) bool {
	if len(arr1) < len(arr2) {
		return false
	}

	for i, el2 := range arr2 {
		if arr1[i] != el2 {
			return false
		}
	}
	return true
}
