// https://go-tour-jp.appspot.com/moretypes/19
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// mapリテラルはstructリテラルに似ているがキーが必要
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
// https://go-tour-jp.appspot.com/moretypes/20
