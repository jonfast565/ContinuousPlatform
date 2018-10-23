package stringutil

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
