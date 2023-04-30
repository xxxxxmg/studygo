package main

import "fmt"

func main() {
	str := "hello"
	slice := []byte(str)
	fmt.Println(slice)

	for _, v := range slice {
		fmt.Println(v)
	}
}
