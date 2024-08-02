// 需求：
// 要求统计  1-200000    的数字中，哪些是素数？这个问题在本章开篇就提出了，现在我们有  goroutine
// 和  channel  的知识后，就可以完成了    [测试数据: 80000]

package main

import "fmt"

// 向 intChan 放入 1-8000 个数
func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}
	close(intChan)
}

// 从 intChan  取出数据，并判断是否为素数,如果是，就放入到  primeChan
func  primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个 primeNum 协程因为取不到数据，退出")
	exitChan <- true
}


func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 4)

	go putNum(intChan)

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%v\n", res)
	}
	fmt.Println("main 线程退出")
}