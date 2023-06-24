package main

import "fmt"

func main() {

	s := Num("1234554321")
	fmt.Println(s)
}

func Num(str string) bool {
	s := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			s = false
		}
	}
	return s
}
