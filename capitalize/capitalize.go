package goreloaded

import "strings"

func Capitalize(word string) string {
	word = strings.ToUpper(word[:1]) + word[1:]
	return word
}
