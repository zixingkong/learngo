package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Person struct {
	Name string
	Age int
	Address string
}


func (p Person)String() string {
	return "Name:" + p.Name + " ,Age:" + strconv.Itoa(p.Age) + " ,Address:" + p.Address
}

func main(){
	
	personChan := make(chan interface{}, 10)

	for i := 0; i < 10; i++ {
		person := Person{
			Name: "name"+strconv.Itoa(i),
			Age: rand.Intn(100),
			Address: "address",
		}
		personChan <-person
	}

	for i := 0; i < 10; i++ {
		person := <-personChan
		fmt.Println(person)
	}

}