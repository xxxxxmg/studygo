package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func ReaderFile(id int, wg *sync.WaitGroup) string {
	defer wg.Done()
	name := fmt.Sprintf("%d.log", id)
	path := "./" + name
	context, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("open file err", err)
	}
	str := string(context)
	return str
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	file, err := os.Create("./hello.log")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var ch = make(chan string)
	//str := make([]string, 0)
	str := ""
	var i int
	for i = 0; i < 10; i++ {
		go func(a int) { //这里如果没有值传递，就会因为for循环太快，创建线程时间跟不上，就会出现 系统找不到 10.log
			str := ReaderFile(a, &wg)
			ch <- str
		}(i)
	}
	for i := 0; i < 10; i++ {
		s := <-ch
		str += s
	}
	wg.Wait()

	ss := strings.Split(strings.TrimSpace(str), "\n")

	countext := make([]int, 0, len(ss))
	for _, v := range ss {
		i, _ := strconv.Atoi(v)
		countext = append(countext, i)
	}

	sort.Ints(countext)
	count := 0
	for i := 0; i < len(countext); i++ {
		count++
		str := fmt.Sprintf("%d ", countext[i])
		if count == 10 {
			str = fmt.Sprintf("%d\n", countext[i])
			count = 0
		}
		file.WriteString(str)
	}

}

//left表示数组左边的下标
//right表示数组右边的下标
func QuickSort(left int, right int, array []int64) {
	l := left
	r := right
	// pivot 表示中轴
	pivot := array[(left+right)/2]
	//for循环的目标是将比pivot小的数放到左边，比pivot大的数放到右边
	for l < r {
		//从pivot左边找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		//从pivot右边找到大于等于pivot的值
		for array[r] > pivot {
			r--
		}
		//交换位置
		array[l], array[r] = array[r], array[l]
		//优化
		if l == r {
			l++
			r--
		}
		//向左递归
		if left < r {
			QuickSort(left, r, array)
		}
		//向左递归
		if right > l {
			QuickSort(l, right, array)
		}
	}
}
