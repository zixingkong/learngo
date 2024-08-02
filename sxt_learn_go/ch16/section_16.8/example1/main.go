// 请完成goroutine和channel协同工作的案例，具体要求：
// 1. 开启一个writeData，向管道intChan中放入50个整数
// 2. 开启一个readData，从intChan中读取writeData放入的数据
// 3. 注意：writeData和readData操作的是同一个管道
// 4. 主线程需要等待writeData和readData协程都完成后才能退出

// channel  支持  for--range  的方式进行遍历，请注意两个细节
// 1)   在遍历时，如果  channel  没有关闭，则回出现  deadlock  的错误
// 2)   在遍历时，如果  channel  已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("writeData:", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}


func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	// go readData(intChan, exitChan)
	// for {
	// 	_, ok := <-exitChan
	// 	if !ok {
	// 		break
	// 	}
	// }
	for _ = range exitChan{
		fmt.Print("OK")
	}
}