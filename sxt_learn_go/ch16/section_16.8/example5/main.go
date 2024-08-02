package main

import (
	"fmt"
	"time"
)

// goroutine中使用recover，解决协程中的panic，阻止程序崩溃，程序可以继续执行
func sysHello()  {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

func  test()  {
	defer func ()  {
		if err:=recover();err!=nil{
			fmt.Println("test()发生错误",err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang"
}

func main()  {
	go sysHello()
	go test()


	for i:=0;i<10;i++{
		fmt.Println("main() ok=",i)
		time.Sleep(time.Second)
	}
}