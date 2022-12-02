package main

import (
	"fmt"
	"strings"
)

func main() {
	input := `199
200
208
210
200
207
240
269
260
263`

	arr := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	for _, s := range arr {
		fmt.Println("Hello,", s)
	}
}
