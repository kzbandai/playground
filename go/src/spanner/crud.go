package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/spanner"
)

type st struct {
}

const (
	gcpProjectID        = ""
	spannerInstanceName = ""
	databaseName        = ""

	tableName = ""
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c, err := newSpannerClient(ctx)
	if err != nil {
		exit(err)
	}

	s := st{}

	m, err := spanner.InsertOrUpdateStruct(tableName, s)
	_, err = c.Apply(ctx, []*spanner.Mutation{m})
	if err != nil {
		exit(err)
	}

}

func exit(err error) {
	fmt.Print(err)
	os.Exit(1)
}

func newSpannerClient(ctx context.Context) (*spanner.Client, error) {
	dbn := "projects/" + gcpProjectID + "/instances/" + spannerInstanceName + "/databases/" + databaseName
	client, err := spanner.NewClient(ctx, dbn)
	if err != nil {
		return nil, err
	}

	return client, nil
}
