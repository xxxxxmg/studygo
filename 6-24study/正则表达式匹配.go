package main

import (
	"fmt"
	"strings"
)

func main() {

	s := Str("abcdf", "a.*")
	fmt.Println(s)

}

func Str(s, p string) bool {
	//判断是否完全相等 或者p字符串是否是“.*”
	if s == p || p == ".*" {
		return true
	}

	p1 := strings.Split(p, "")
	s1 := strings.Split(s, "")

	//判断p字符串是否包含*号
	p1j := 0
	for i := 0; i < len(p1); i++ {
		if p1[i] == "*" {
			p1j = i
			break
		}
	}

	if p1j == 0 && len(p1) != len(s1) {
		return false
	}

	//字符串包含*号
	if p1j != 0 {
		for i := 0; i < p1j; i++ {
			if s1[i] == p1[i] || p1[i] == "." {
				continue
			}
			return false
		}
		return true
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] == p1[i] || p1[i] == "." {
			continue
		}
		return false
	}

	return true
}
