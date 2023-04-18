package main

import "fmt"

type Student struct {
	Name string
	age  int
}

type Stu interface{} //空接口

func main() {
	var s Stu
	stud := Student{"李四", 24}
	s = stud // 空接口可以接任何类型的数据，但是不能做运算
	fmt.Println(s)
	var i1 Stu
	var j1 Stu
	i1 = 10
	j1 = 20
	// sum := i1 + j1 空接口可以接任何类型的数据，但是不能做运算
	fmt.Println(i1)
	fmt.Printf("%T\n", i1)
	fmt.Printf("%T\n", j1)

	var slice []interface{} //定义了一个空接口类型的切片

	slice = append(slice, 1, 2, 3, [3]int{4, 5, 6}, true, "hello")
	fmt.Println(slice...)

	//通过类型断言，获取数据类型
	//数据的值，OK := 空接口数据.(数据类型)
	value, ok := slice[0].(int)
	//还有 通过反射获取空接口数据reflect
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("0号索引的值不是int类型的数据")
	}
}
