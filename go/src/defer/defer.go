package main

import "fmt"

func LearnDefer() {
	defer fmt.Println("World")
	fmt.Print("Hello")
}

func LearnDefer2() {
	// deferされたものは関数の脱出直前に実行される
	fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("5")
	defer fmt.Println("6")

	/*
		1
		5
		6
		4
		3
		2
	*/
}

func main() {
	LearnDefer() // HelloWorld
	LearnDefer2()
}
