package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"io"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	opts := option.ClientOption(
		option.WithCredentialsFile(os.Getenv("CredentialsFile")),
	)

	client, err := storage.NewClient(ctx, opts)
	if err != nil {
		log.Fatalf("%v", err)
	}
	filename := "mmczblsq.doc"
	filepath := fmt.Sprintf("./tmp/%s", filename)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()

	uuidIns, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("%v", err)
	}
	object := fmt.Sprintf("%s/%s", uuidIns, filename)
	log.Printf("object name: %s", object)
	wc := client.Bucket("ez2on").Object(object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		log.Fatalf("%v", err)
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("%v", err)
	}
}
