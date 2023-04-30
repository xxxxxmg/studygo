package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建一个新文件
	// fp, err := os.Create("./hello.txt")
	// defer fp.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//以追加字符串的方式，打开一个文件
	fp, err := os.OpenFile("./hello2.txt", os.O_APPEND, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	//要追加的字符串内容
	n, err := fp.WriteString(" hello world")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功:", n)
}
