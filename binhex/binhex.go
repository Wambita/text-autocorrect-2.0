package goreloaded

import (
	"fmt"
	"strconv"
)

func BinToDec(binStr string) string {
	decStr, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.FormatInt(decStr, 10)
}

func HexToDec(hexStr string) string {
	decStr, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.FormatInt(decStr, 10)
}
