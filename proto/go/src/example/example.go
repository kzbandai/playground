package main

import (
	"fmt"
	"strconv"
	//	"foo_bar"
	//	"github.com/golang/protobuf/proto"
	//	"io/ioutil"
)

func main() {
	var stdin string
	in, err := fmt.Scan(&stdin)
	if err != nil {
		fmt.Printf("error!")
	}

	fmt.Printf(strconv.Itoa(in))
}
