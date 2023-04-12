package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		// 注释内容也可以使用
		// idleDuration := time.NewTimer(time.Second * 10).C
		idleDuration := time.After(time.Second * 10)
		ticker := time.NewTicker(time.Second * 1)

		for {
			select {
			case <-ticker.C:
				fmt.Println("hello world")
			case <-idleDuration:
				fmt.Println("exit")
				return
			}
		}
	}()

	time.Sleep(time.Second * 15)
	fmt.Println("main func")
}
