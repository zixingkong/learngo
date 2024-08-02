package main

import "fmt"

// channel只读和只写的最佳实践

// ch chan <- int 只能写入
func send(ch chan<-int,exitChan chan struct{}){
	for i:=0;i<10;i++{
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}

// ch <- chan int 只能读取
func recv(ch <- chan int,exitChan chan struct{})  {
	for {
		v,ok := <-ch
		if !ok{
			break
		}
		println(v)
	}
	var a struct{}
	exitChan <- a
}

func main()  {
	ch := make(chan int,10)
	exitChan := make(chan struct{},2)

	go send(ch,exitChan)
	go recv(ch,exitChan)

	var total = 0 
	for _ = range exitChan{
		total++
		if total == 2{
			break
		}
	}
	fmt.Println("结束。。。")
}