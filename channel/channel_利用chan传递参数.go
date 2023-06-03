package main

import "fmt"

var ch = make(chan int)

func sprint(i1, i2 int) int {
	return i1 + i2
}

func user1() {
	v := sprint(2, 3)
	ch <- v
}
func user2() {
	v := <-ch
	fmt.Println("v+5的结果为：", v+5)
}

func main() {

	go user1()
	go user2()
	for {
	}
}
