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
	}

	// up, low and cap modifications
	words := strings.Fields(line)
	for i, word := range words {
		if word == "(cap)" || word == "(CAP)" || word == "(Cap)" {
			words[i-1] = r.Capitalize(words[i-1])
			words = append(words[:1], words[i+1:]...)
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
			}
			words = append(words[:i], words[i+1:]...)
		} else if word == "(low)" || word == "(LOW)" || word == "(Low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(low," || word == "(LOW," || word == "(Low," {
			numStr := strings.Trim(words[i+2], ")")
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
			}
			words = append(words[:i], words[i+2:]...)
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
			}
			words = append(words[:i], words[i+2:]...)
		}
	}

	return strings.Join(words, " ")
}
