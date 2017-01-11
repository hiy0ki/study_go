package main

// factored import statement という書き方 グループ化してまとめて
import (
    "fmt"
    "math/rand"
    "time"
)

// import "fmt"
// import "time
// このようにも書けるが、複数ある場合はfoctoredな書き方を推奨されている


func main() {
    num := time.Now().UnixNano()
    fmt.Println(num)
    rand.Seed(num)

    // export(公開)されたものの名前は大文字にするルール  TODO 命名規則なのか一応確認
    fmt.Println("my favorite number is ", rand.Intn(10))
}
