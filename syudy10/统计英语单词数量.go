package main

import (
	"fmt"
	"strings"
)

//统计有多少个单词
func main() {
	//英文字符串
	str := "as an & operand to any format that accepts a string"
	//按照空格分割
	slice := strings.Split(str, " ")
	num := 0
	for _, v := range slice {
		s := true
		for _, ch := range v {
			if !(ch >= 65 && ch <= 90 || ch >= 97 && ch <= 122) {
				s = false
				break
			}
		}
		if s {
			num++
			fmt.Println(v)
		}

	}

	//输出单词数组的长度
	fmt.Println(num)

}
