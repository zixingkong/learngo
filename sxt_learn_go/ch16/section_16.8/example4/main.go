package main

// 使用select可以解决从管道读取数据阻塞的问题

import (
	"fmt"
)


func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时，如果不关闭会阻塞而导致deadlock
	// label:
	for {
		select {
		// 注意：这里如果intChan一直没有关闭，会造成阻塞而导致deadlock
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
		default:
			fmt.Println("intChan和stringChan均没有数据")
			// break label
			return
		}
	}
}