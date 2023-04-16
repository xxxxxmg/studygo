package main

import (
	"fmt"

	study "studay02/sub"

	_ "github.com/gin-gonic/gin"
	"github.com/mht/sutdy01/add"
	"github.com/mht/sutdy01/add/add3"
	"github.com/mht/sutdy01/chenf"
	cfAdd "github.com/mht/sutdy01/chenf/add3"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(add.Add(1, 2))
	fmt.Println(study.Sub(1, 2))
	fmt.Println(chenf.Chenf(10, 19))
	fmt.Println(add3.Add(8, 2))
	fmt.Println(cfAdd.Add(8, 2))
	f := add2(10)       // f == 10
	fmt.Println(f(5))   // 15
	fmt.Println(f(5))   // 20
	fmt.Println(f(100)) // 120

	fmt.Println("icnr ----- ")
	in := incr(100)
	for i := 0; i < 10; i++ {
		fmt.Println(in())
	}
	fmt.Println(compute(func(f1, f2 float64) float64 {
		return f1 * f2
	}))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func incr(a int) func() int {
	return func() int {
		a += 10
		return a
	}
}

// a == 10
func add2(a int) func(int) int {
	return func(i int) int {
		a += i // 15 += 5 == 20
		return a
	}
}
