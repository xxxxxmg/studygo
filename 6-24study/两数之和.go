package main

import "fmt"

func main() {

	var nums = []int{2, 7, 11, 15}
	Num(nums, 19)

}

func Num(num []int, target int) {

	for i := 0; i < len(num); i++ {
		for j := 1; j < len(num)-1; j++ {
			if num[i]+num[j] == target {
				fmt.Println(target, "=", num[i], "+", num[j], "下标为", i, j)
				return
			}
		}
	}
	fmt.Println("数组中两数相加没有结果为", target, "的数字")
}
