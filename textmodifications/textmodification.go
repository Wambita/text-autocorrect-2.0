package goreloaded

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	goreloaded "goreloaded/binhex"
	r "goreloaded/capitalize"
)

func Textmodification(line string) string {
	// binary
	line = binaryModification(line)
	// hex
	line = hexadecimalModification(line)
	// cap
	line = modifyCase(line, "cap")
	// low
	line = modifyCase(line, "low")
	// up
	line = modifyCase(line, "up")
	// vowel
	line = vowelModification(line)
	// punc
	line = punctuationModification(line)
	return line
}

// general function for case transformation (up), low and cap
func modifyCase(line, action string) string {
	words := strings.Fields(line)
	for i, word := range words {
		if strings.HasPrefix(word, "("+action+")") {
			// remove the action from the word
			words = append(words[:i], words[i+1:]...)
			// apply the transformation
			words[i-1] = transformWord(words[i-1], action)
		} else if strings.HasPrefix(word, "("+action+",") {
			// number based transformations
			numStr := strings.Trim(words[i+1], ")")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println(err)
			}
			if num > len(words[i-1]) {
				fmt.Println("Error: Number larger than available words to transform")
			} else {
				for j := 0; j <= num; j++ {
					words[i-j] = transformWord(words[i-j], action)
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
	}
	return strings.Join(words, " ")
}

// helper function for applying the transformation based on the action
func transformWord(word, action string) string {
	switch action {
	case "cap":
		return r.Capitalize(word)
	case "low":
		return strings.ToLower(word)
	case "up":
		return strings.ToUpper(word)
	}
	return word
}

// case transformations
// cap

// binary to  decimal
func binaryModification(line string) string {
	return modifyNum(line, `(?i)([0-1]+)\s?\(bin\)`, "(bin)", goreloaded.BinToDec)
}

// hexaecimal to decimal
func hexadecimalModification(line string) string {
	return modifyNum(line, `([0-9A-Fa-f]+)\s?\(hex\)`, "(hex)", goreloaded.HexToDec)
}

// helper function to modify number based on regex pattern
func modifyNum(line, pattern, replacement string, convertFunc func(string) string) string {
	matches := regexp.MustCompile(pattern).FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		patt := match[1]
		decStr := convertFunc(patt)
		line = strings.Replace(line, patt+" "+replacement, decStr, 1)
	}
	return line
}

// vowel modification
func vowelModification(line string) string {
	words := strings.Fields(line)
	vowels := "aeiouhAEIOUH"

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			for _, v := range vowels {
				if rune(words[i+1][0]) == v {
					words[i] = modifyVowels(words[i])
				}
			}
		}
	}
	return strings.Join(words, " ")
}

// modify vowels
func modifyVowels(word string) string {
	if word == "a" {
		return "an"
	} else if word == "A" {
		return "An"
	}
	return word
}

// punctuation modification
func punctuationModification(line string) string {
	// list of common punc patterns
	punctuations := []struct {
		pattern     string
		replacement string
	}{
		{` '`, "'"},                           // Remove spaces before single quotes
		{`' `, "'"},                           // Remove spaces after single quotes
		{`(.*?)(\s+)([.,!?:;])`, "$1$3 "},     // Remove spaces between text and punctuation
		{`([.,!?:;])(\s+)([.,!?:;])`, "$1$3"}, // Handle multiple punctuation marks
		{`(\s+)(.*?)`, " $2"},                 // Handle leading spaces
		{`(:)(')(.*?)`, "$1 $2"},              // Handle cases like :'char
		{`(\.)(\s)(')`, "$1$3"},               // Handle cases like . 'char
		{`([?])(\s+)`, "$1"},                  // Remove spaces after a question mark

	}

	// apply punc modifications
	for _, punc := range punctuations {
		line = regexp.MustCompile(punc.pattern).ReplaceAllString(line, punc.replacement)
	}
	return line
}
