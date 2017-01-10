package main

import (
    "fmt"
    "math/rand"
    "time"
)


func main() {
    num := time.Now().UnixNano()
    fmt.Println(num)
    rand.Seed(num)
    fmt.Println("my favorite number is ", rand.Intn(10))
}
