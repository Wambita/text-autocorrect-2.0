package goreloaded

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	reload "goreloaded/binhex"
	r "goreloaded/capitalize"
)

func binaryModifcation(line string) string {
	// binary to decimal modification
	binmatches := regexp.MustCompile(`(?i)([0-1]+)\s?\(bin)`).FindAllStringSubmatch(line, -1)
	for _, match := range binmatches {
		line = strings.Replace(line, match[1]+" (bin)", reload.BinToDec(match[1]), 1)
	}
	return line
}

func hexadecimalModifcation(line string) string {
	// hexadecimal to decimal modification
	hexmatches := regexp.MustCompile(`(?i)([0-9A-Fa-f]+)\s?\(hex)`).FindAllStringSubmatch(line, -1)
	for _, match := range hexmatches {
		line = strings.Replace(line, match[1]+" (hex)", reload.HexToDec(match[1]), 1)
	}
	return line
}

// capitalize
func capitalize(line string) string {
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

		}
	}
	line = strings.Join(words, " ")
	return line
}

// lowercase
func lowerCase(line string) string {
	words := strings.Fields(line)
	for i, word := range words {
		if word == "(low)" || word == "(LOW)" || word == "(Low)" {
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

		}
	}
	line = strings.Join(words, " ")
	return line
}

// upper case
func upperCase(line string) string {
	words := strings.Fields(line)
	for i, word := range words {
		if word == "(up)" || word == "(UP)" || word == "(Up)" {
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
	line = strings.Join(words, " ")
	return line
}

// vowels
func vowelModification(line string) string {
	words := strings.Fields(line)
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
	line = strings.Join(words, " ")
	return line

}


// punctuation modification
func punctuationModification(line string) string{
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



func TextModification(line string) string {
	// Hexadecimal to decimal modification
	
	return line
}
