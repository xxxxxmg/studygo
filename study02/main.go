package main

import (
	"fmt"
	"os"
)

type Calcer interface {
	Result() int
}

type Operate struct {
	a1 int
	a2 int
}

type Add struct {
	Operate
}
type Sub struct {
	Operate
}

func (a *Add) Result() int {
	return a.a1 + a.a2
}

func (s *Sub) Result() int {
	return s.a1 - s.a2
}

type Student struct {
	name string
}

func (s *Student) Write(p []byte) (n int, err error) {
	fmt.Println(s.name, string(p))
	return 0, nil
}

func main() {
	//定义接口变量
	var c Calcer

	//实例化加法对象
	//a := Add{Operate{10, 20}}
	//将对象赋值给接口
	//c = &a

	//实例化减法对象
	s := Sub{Operate{30, 10}}
	c = &s

	// O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// // The remaining values may be or'ed in to control behavior.
	// O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	// O_CREATE int = syscall.O_CREAT  // create a new file if none exists.

	//获取返回值
	vlaue := c.Result()
	fmt.Println(vlaue)
	stu := Student{name: "张三"}
	fmt.Fprint(&stu, "ssssss")

	mht, err := os.OpenFile("stu.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	os.Stdout = mht
	fmt.Println("ssssssssssss")

}
