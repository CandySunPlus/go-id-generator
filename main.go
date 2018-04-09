package main

import (
	"fmt"

	"sfmblog.cn/idgenEx/idgen"
)

var channel chan int = make(chan int)

func loop(idGenerator *idgen.IDGenerator) {
	for i := 0; i < 100; i++ {
		id, err := idGenerator.Next("test", 12345678)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v -> %d \n", idgen.ParseID(id), id)
	}
	channel <- 0
}

func main() {
	idGenerator := idgen.NewIDGenerator("119.27.180.231", 6379, "c5809078fa6d652e0b0232d552a9d06d37fe819c")
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
