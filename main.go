package main

import (
	"fmt"

	IdGen "./idgen"
)

var channel chan int = make(chan int)

func loop(idGenerator *IdGen.IDGenerator) {
	for i := 0; i < 100; i++ {
		id, err := idGenerator.Next("test", 12345678)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v -> %d \n", IdGen.ParseID(id), id)
	}
	channel <- 0
}

func main() {
	idGenerator := IdGen.NewIDGenerator("119.27.180.231", 6379, "c5809078fa6d652e0b0232d552a9d06d37fe819c")
	go loop(idGenerator)
	go loop(idGenerator)
	go loop(idGenerator)
	go loop(idGenerator)
	go loop(idGenerator)
	go loop(idGenerator)
	go loop(idGenerator)
	go loop(idGenerator)
	<-channel
}
