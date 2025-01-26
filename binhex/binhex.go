package goreloaded

import (
	"fmt"
	"strconv"
)

//convert hexadecimal string to decimal string
func BinToDec(binStr string) string {
	decStr, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.FormatInt(decStr, 10)
}

//convert binary string to decimal string
func HexToDec(hexStr string) string {
	decStr, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.FormatInt(decStr, 10)
}
