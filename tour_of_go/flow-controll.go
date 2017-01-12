package main

import "fmt"

func main(){

    // 繰り返し制御はfor文しかない
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    sum2 := 1
    for ; sum2 < 10000; {
        sum2 += sum2
    }
    fmt.Println(sum2)
}
