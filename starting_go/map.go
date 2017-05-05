package main

import (
	"fmt"
)

func main() {
	// 宣言
	//var m map[int]string

	// makeで作る
	m := make(map[int]string)

	m[1] = "hogehoge"
	m[30] = "hugahuga"
	m[-1] = "hahahah"

	fmt.Println(m)

	// mapリテラルで作る
	m2 := map[string]string{
		"us":   "america",
		"jp":   "japan",
		"hoge": "hogehoge", // カンマが必要
	}

	fmt.Println(m2)

	// sliceをもつmap
	m3 := map[int][]int{
		1: []int{1},
		2: []int{1, 3},
		3: {4, 3, 3}, // 省略することも可能
	}

	fmt.Println(m3)

	// mapの要素がmapの場合 宣言が複雑なのは仕方ないのか。。
	m4 := map[int]map[float64]string{
		1: {3.14: "円周率"},
	}

	fmt.Println(m4)

	// 要素の参照時にキーがない場合に初期値が返されてしまう。
	// その回避方法
	m5 := map[int]string{1: "a", 2: "b", 3: "c"}
	s, ok := m5[1]
	fmt.Println(s, ok)
	s2, ok2 := m5[5]
	fmt.Println(s2, ok2)

	// よくある書き方
	if _, ok := m5[2]; ok {
		// キーが存在する場合に処理
		fmt.Println("キーが存在するよ")
	}

	// forで回す
	// キーの順序は保証されない
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// 要素を削除する
	delete(m5, 2)
	fmt.Println(m5)
}
