package main

import (
	"fmt"
	"strings"
)

func Str(s string) string {

	str2 := strings.Split(s, "_")

	str3 := ""

	for i := 0; i < len(str2); i++ {
		str3 += strings.Title(str2[i])
	}

	return str3
}

func main() {
	s := Str("User_login_2023_6_19")
	fmt.Println(s)
}
