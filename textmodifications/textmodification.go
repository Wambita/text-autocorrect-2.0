package goreloaded

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	reload "goreloaded/binhex"
	r "goreloaded/capitalize"
)

func TextModification(line string) string {
	// binary to decimal modification
	binMatches := regexp.MustCompile(`([01]+)\s?\(bin\)`).FindAllStringSubmatch(line, -1)
	for _, match := range binMatches {
		binStr := match[1]
		decStr := reload.BinToDec(binStr)
		line = strings.Replace(line, binStr+" (bin)", decStr, 1)
	}

	// Hexadecimal to decimal modification
	hexMatches := regexp.MustCompile(`([0-9A-Fa-f]+)\s?\(hex\)`).FindAllStringSubmatch(line, -1)
	for _, match := range hexMatches {
		hexStr := match[1]
		decStr := reload.HexToDec(hexStr)
		line = strings.Replace(line, hexStr+" (hex)", decStr, 1)
		line = strings.Replace(line, hexStr+" (hex)", decStr, 1)
	}

	// up, low and cap modifications
	words := strings.Fields(line)
	for i, word := range words {
		if word == "(cap)" || word == "(CAP)" || word == "(Cap)" {
			words[i-1] = r.Capitalize(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(cap," || word == "(CAP" || word == "(Cap" {
			numStr := strings.Trim(words[i+1], ")")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println(err)
			}
			if num > len(words[i-1]) {
				fmt.Println("Error: Number larger than available words to capitalize")
			} else {
				for j := 0; j <= num; j++ {
					words[i-j] = r.Capitalize(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			}

		} else if word == "(low)" || word == "(LOW)" || word == "(Low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(low," || word == "(LOW," || word == "(Low," {
			numStr := strings.Trim(words[i+1], ")")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println(err)
			}
			if num > len(words[i-1]) {
				fmt.Println("Error: Number larger than available words to change to lowercase")
			} else {
				for j := 0; j <= num; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			}
		} else if word == "(up)" || word == "(UP)" || word == "(Up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(up," || word == "(UP," || word == "(Up," {
			numStr := strings.Trim(words[i+1], ")")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println(err)
			}
			if num > len(words[i-1]) {
				fmt.Println("Error: Number larger than available words to change to uppercase")
			} else {
				for j := 0; j <= num; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			}

		}
	}

	// vowels correction
	for i := 0; i < len(words)-1; i++ {
		// Vowel correction for a,e,i,o,u and h capital and small
		vowels := "aeiouhAEIOUH"
		for j := 0; j < len(vowels)-1; j++ {
			if words[i] == "a" && words[i+1][0] == vowels[j] {
				words[i] = "an"
			} else if words[i] == "A" && words[i+1][0] == vowels[j] {
				words[i] = "An"
			}
		}
	}
	line = strings.Join(words, " ") // convert to string
	// '' punctuation , removing spaces
	line = strings.ReplaceAll(line, " '", "'")
	line = strings.ReplaceAll(line, "' ", "'")
	// Punctuation modification for this combination "my ," to "my,"
	punctuate := `(.*?)(\s+)([.,!?:;])` // all characters plus space(s)plus punctuation marks
	newStr := regexp.MustCompile(punctuate)
	line = newStr.ReplaceAllString(line, "$1$3 ")
	// Punctuation modification for combination my . . . to my...
	punctuate2 := `([.,!?:;])(\s+)([.,!?:;])` // punctuation marks plus spaces plus punctuation marks
	newStr2 := regexp.MustCompile(punctuate2)
	line = newStr2.ReplaceAllString(line, "$1$3")
	// punctuation marks plus spaces plus punctuation marks
	punctuate3 := `(\s+)(.*?)`
	newStr3 := regexp.MustCompile(punctuate3)
	line = newStr3.ReplaceAllString(line, " $2")
	// // case of :'char
	punctuate4 := `(:)(')(.*?)`
	newStr4 := regexp.MustCompile(punctuate4)
	line = newStr4.ReplaceAllString(line, "$1 $2")
	// case 5
	punctuate5 := `(\.)(\s)(')`
	newStr5 := regexp.MustCompile(punctuate5)
	line = newStr5.ReplaceAllString(line, "$1$3")

	// case 6
	punctuate6 := `([?])(\s+)`
	newStr6 := regexp.MustCompile(punctuate6)
	line = newStr6.ReplaceAllString(line, "$1")
	return line
	// join modified words back into single line
}
