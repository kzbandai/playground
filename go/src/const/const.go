package main

import "fmt"

const (
	OUTER = "外側のconst"
)

func main() {
	const inner = "内側のconst"
	fmt.Printf("%s %s\n", OUTER, inner)
}
