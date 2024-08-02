// 需求：现在要计算    1-200    的各个数的阶乘，并且把各个数的阶乘放入到  map  中。 最后显示出来。要求使用  goroutine  完成

package main

import "time"

var myMap = make(map[int]int, 10)




func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}

func main()  {
	for i:=1; i<=20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	for i, v := range myMap {
		println(i, "的阶乘是：", v)
	}
}