package stringutil

import "regexp"

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
