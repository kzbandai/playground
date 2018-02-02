package main

import "fmt"

const (
	A = iota
	B
	C
	_
	D
	E
)

// コンテキストが異なるとiotaは0に戻る
// 同じコンテキスト内では値が1ずつ増加する

func main() {
	const X = iota
	const Y = iota

	fmt.Printf("%d\n", A) // 0
	fmt.Printf("%d\n", B) // 1
	fmt.Printf("%d\n", C) // 2
	fmt.Printf("%d\n", D) // 4
	fmt.Printf("%d\n", E) // 5

	fmt.Printf("%d\n", X) // 0
	fmt.Printf("%d\n", Y) // 0
}
