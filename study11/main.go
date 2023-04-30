package main

import "fmt"

func main() {
	// 初始化一个空的学生map
	stu := make(map[int]map[string][]int)

	// 添加学生数据
	stu[101] = map[string][]int{"yuwen": {80, 70}, "shuxue": {90, 80}}
	stu[102] = map[string][]int{"yuwen": {85, 83}, "shuxue": {95, 98}}

	// 遍历学生map，输出每个学生的信息
	for id, _ := range stu {
		// fmt.Printf("ID：%d\n", id)
		// for subject, grades := range scores {
		// 	fmt.Printf("%s成绩：%v\n", subject, grades)
		// }

		// 将新的数据添加到嵌套的map中
		stu[id]["yingyu"] = []int{75, 85}
	}

	// 再次遍历学生map，输出每个学生的信息
	for id, scores := range stu {
		fmt.Printf("ID：%d\n", id)
		for subject, grades := range scores {
			fmt.Printf("%s成绩：%v\n", subject, grades)
		}
	}
}
