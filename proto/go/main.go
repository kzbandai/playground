package main

import (
	"fmt"
	"strconv"
)

func main() {
	var stdin string
	in, err := fmt.Scan(&stdin)
	if err != nil {
		fmt.Printf("error!")
	}

	fmt.Printf(strconv.Itoa(in))
}
