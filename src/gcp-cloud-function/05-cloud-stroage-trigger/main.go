package cloudstoragetrigger

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"gcp-cloud-function/05-cloud-stroage-trigger/utils/slice"
	"os"
	"time"

	"cloud.google.com/go/functions/metadata"
)

type GCSEvent struct {
	Bucket         string    `json:"bucket"`
	Name           string    `json:"name"`
	Metageneration string    `json:"metageneration"`
	ResourceState  string    `json:"resourceState"`
	TimeCreated    time.Time `json:"timeCreated"`
	Updated        time.Time `json:"updated"`
}

var (
	bucketName           = os.Getenv("bucket")
	imageMimeTypes       = []string{"image/gif", "image/png", "image/jpeg", "image/bmp", "image/webp"}
	videoMimeTypes       = []string{"video/mp4"}
	mimeTypes            = append(imageMimeTypes, videoMimeTypes...)
	MaxSize        int64 = 1024 * 1
)

func CloudStorageTriggerTest(ctx context.Context, e GCSEvent) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	dstObject := client.Bucket(bucketName).Object(e.Name)
	_, err = dstObject.Attrs(ctx)
	if err == nil {
		fmt.Printf("object: %s has already existed", e.Name)
		return nil
	}
	if err != storage.ErrObjectNotExist {
		return err
	}
	srcObjectName := fmt.Sprintf("/tmp/%s", e.Name)
	srcObject := client.Bucket(bucketName).Object(srcObjectName)
	if err := validate(ctx, srcObject); err != nil {
		return err
	}
	return nil
}

func validate(ctx context.Context, object *storage.ObjectHandle) error {
	objectAttrs, err := object.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("failed to get object attributes: %q", object.Name)
	}

	if objectAttrs.Size > MaxSize {
		return fmt.Errorf("file is larger than %d, size: %d", MaxSize, objectAttrs.Size)
	}

	slice.Contains(mimeTypes, objectAttrs.ContentType)

	return nil
}

func validateMIMEType(ctx context.Context, attrs *storage.ObjectAttrs, obj *storage.ObjectHandle) error {
	r, err := obj.NewReader(ctx)
	if err != nil {
		return fmt.Errorf("failed to open file %q", obj.ObjectName())
	}
	defer r.Close()

	return nil
}

func printMetaData(ctx context.Context) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	fmt.Printf("Event ID: %v\n", meta.EventID)
	fmt.Printf("Event type: %v\n", meta.EventType)
	fmt.Printf("Bucket: %v\n", e.Bucket)
	fmt.Printf("File: %v\n", e.Name)
	fmt.Printf("Metageneration: %v\n", e.Metageneration)
	fmt.Printf("Created: %v\n", e.TimeCreated)
	fmt.Printf("Updated: %v\n", e.Updated)
	return nil
}
