package main

import "fmt"

func main() {
	ch1 := make(chan int) //ch1通道传输数据

	ch2 := make(chan int) //协调主程序和子程序的执行顺序  先写入，在读取

	go func() {
		for i := 0; i < 30; i++ {
			ch1 <- i //将i写入到ch1通道中
			fmt.Printf("子进程%d写到主程序\n", i)
			ch2 <- 0 //将0写入到ch2通道中
		}
	}()

	for i := 0; i < 30; i++ {
		num := <-ch1 //读取ch1通道中的数据，并赋值给num
		<-ch2        //读取ch2通道，并丢弃数据	如果子程序没有将数据写到本通道，这里会形成阻塞，直到本通道中有数据写入
		fmt.Println("主程序读到：", num)
	}

}
