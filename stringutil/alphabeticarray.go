package stringutil

import "strings"

type AlphabeticArray []string

func (list AlphabeticArray) Len() int { return len(list) }

func (list AlphabeticArray) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list AlphabeticArray) Less(i, j int) bool {
	var si = list[i]
	var sj = list[j]
	var si_lower = strings.ToLower(si)
	var sj_lower = strings.ToLower(sj)
	if si_lower == sj_lower {
		return si < sj
	}
	return si_lower < sj_lower
}
