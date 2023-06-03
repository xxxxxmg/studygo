package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go test(i, &wg)
	}
	wg.Wait()
}

func test(id int, wg *sync.WaitGroup) {
	name := fmt.Sprintf("%d.log", id)
	fp, err := os.Create("./" + name)
	if err != nil {
		fmt.Println("Creat file err", err)
	}
	defer fp.Close()
	defer wg.Done()
	for i := (id * 10); i < (id*10)+10; i++ {
		if i == (id*10)+9 {
			str := fmt.Sprintf("%d", i)
			fp.WriteString(str)
			return
		}
		str := fmt.Sprintf("%d\n", i)
		fp.WriteString(str)
	}
}
