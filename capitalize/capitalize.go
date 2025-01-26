package goreloaded

import "strings"

func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	if len(word) == 1 {
		return strings.ToUpper(word)
	}
	return strings.ToUpper(word[:1]) + word[1:]
}
