package main

import "fmt"

func main(){
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)


	primes := [6]int{2, 3, 5, 7, 11,}
	fmt.Println(primes)

	// slice 配列への参照のようなもの
	var s[]int = primes[1:4]
	fmt.Println(s)

	s2 := primes[2:4]
	fmt.Println(s2)

	names := [4]string{
		"john",
		"paul",
		"george",
		"gopher",
	}
	
	fmt.Println(names)

	x := names[0:2]
	y := names[1:3]
	fmt.Println(x,y)
	
	x[1] = "XXX"
	fmt.Println(x,y)
	fmt.Println(names)


	q := []int{2,3,45,3,6,7,4,}
	fmt.Println(q)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3,false},
		{2, true},
		{3,false},
		{2, true},
		{3,false},
	}
	fmt.Println(st)
	fmt.Println(len(st))
	fmt.Println(cap(st))


	// sliceの初期値はnil
	var aa []int
	fmt.Println(aa,len(aa),cap(aa))
	if aa == nil {
		fmt.Println("nill")
	}

	// make でスライスを作ることができる
	m := make([]int, 5)
	printSlice("m", m)
	
	m1 := make([]int, 0, 5)
	printSlice("m1", m1)

	m2 := m1[:2]
	printSlice("m2", m2)

	m3 := m2[2:5]
	printSlice("m3", m3)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
