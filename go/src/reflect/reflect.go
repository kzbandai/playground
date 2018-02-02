package main

import (
	"fmt"
	"reflect"
)

// http://techblog.raccoon.ne.jp/archives/38280316.html
func main() {
	b := true
	fmt.Println(reflect.TypeOf(b))

	ui8 := uint8(100)
	fmt.Println(reflect.TypeOf(ui8))

	n := 100
	fmt.Println(reflect.TypeOf(n))

	c := 'A'
	fmt.Println(reflect.TypeOf(c))

	r := 'あ'
	fmt.Println(reflect.TypeOf(r))

	u := '\U00101234'
	fmt.Println(reflect.TypeOf(u))

	s := "文字列"
	fmt.Println(reflect.TypeOf(s))

	a := [3]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(a))
}
