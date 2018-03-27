package main

import (
    "fmt"

    IdGen "./idgen"
)

func main() {
    idGenerator := IdGen.NewIDGenerator("119.27.180.231", 6379, "c5809078fa6d652e0b0232d552a9d06d37fe819c")
    id, err := idGenerator.Next("test", 12345678)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(id)
}
