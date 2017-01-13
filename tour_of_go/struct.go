package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X: 1}
	v3 = Vertex{} // X:0,Y:0
	p2  = &Vertex{1,2}
)

func main() {
	fmt.Println(Vertex{1,2})

	v := Vertex{1,2}
	// .(ドット)を用いてアクセスする
	v.X = 4
	fmt.Println(v.X)

	// ポインタを通してアクセスすることもできる
	p := &v
	p.X = 1e9
	fmt.Println(v)


	fmt.Println(v1,v2,v3,p2)
}
