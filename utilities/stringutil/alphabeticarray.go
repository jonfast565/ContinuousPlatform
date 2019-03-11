// Package for dealing with string customizable
package stringutil

import "strings"

// A string array that is alphabetical
type AlphabeticArray []string

// Get the length
func (list AlphabeticArray) Len() int { return len(list) }

// Swap a value
func (list AlphabeticArray) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

// Compare values by using alphabetical values
// NOTE: Method is case insensitive
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
