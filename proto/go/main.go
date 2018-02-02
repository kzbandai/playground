package main

import (
	"fmt"
	p "github.com/golang/protobuf/proto"
	pb "github.com/kzbandai/playground/proto/go/build/pb"
)

func main() {
	sr := &pb.SearchRequest{
		PageNumber: *p.Int32(100),
		Query:      *p.String("hogehoge fugafuga"),
	}

	fmt.Print(sr.String())
}
