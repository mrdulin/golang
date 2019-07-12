package cloudstoragetrigger

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"testing"
)

func TestCloudStorageTriggerTest(t *testing.T) {
	ctx := context.Background()
	opts := option.ClientOption(
		option.WithCredentialsFile("/Users/ldu020/workspace/github.com/mrdulin/nodejs-gcp/.gcp/cloud-storage-admin.json"),
	)
	client, err := storage.NewClient(ctx, opts)
	if err != nil {
		t.Fatalf("storage.NewClient: %v", err)
	}
	bucket := client.Bucket("ez2on")

	it := bucket.Objects(ctx, nil)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Errorf("it.Next() - %v", err)
			return
		}
		fmt.Println(attrs.Name)
	}

	object := bucket.Object("tmp/c26dcec2-e4b5-4fad-8cfe-d73ca39b6b61/ebb9ed02-f1d9-40f7-add8-2ef00391fa4d/Moon.jpg")
	attrs, err := object.Attrs(ctx)
	if err != nil {
		log.Fatalf("object.Attrs(ctx) - %v", err)
	}
	t.Logf("attrs: %v", attrs)
}
