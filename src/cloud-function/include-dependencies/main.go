package includedependenciestest

import (
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
)

const (
	ProjectId = "shadowsocks-218808"
)

func getBucketAttrs()([]*storage.BucketAttrs, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	it := client.Buckets(ctx, ProjectId)

	var bucketAttrsSlice []*storage.BucketAttrs
	for {
		bucketAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		bucketAttrsSlice = append(bucketAttrsSlice, bucketAttrs)
	}
	return bucketAttrsSlice, nil
}

func getBucketNames(bucketAttrsSlice []*storage.BucketAttrs) []string {
	var bucketNames []string
	for _, bucketAttrs := range bucketAttrsSlice {
		bucketNames = append(bucketNames, bucketAttrs.Name)
	}
	return bucketNames
}

func IncludedependenciesTest(w http.ResponseWriter, req *http.Request) {
	bucketAttrsSlice, err := getBucketAttrs()
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, "get bucket attrs error", http.StatusInternalServerError)
		return
	}
	var bucketNames = getBucketNames(bucketAttrsSlice)
	log.Printf("%#v", bucketNames)
	fmt.Fprintf(w, "%#v", bucketNames)
}