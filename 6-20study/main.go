package main

import (
	"fmt"
)

func main() {

	s := "hello golang语言"

	fmt.Println(reverseString(s))

	fmt.Println(reverseString(reverseString(s)))

	// output: 言语gnalog，olleh
	// output: hello，golang语言
}

func reverseString(s string) string {

	runes := []int32(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
