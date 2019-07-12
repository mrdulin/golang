package cloudstoragetrigger

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"gcp-cloud-function/05-cloud-stroage-trigger/utils/slice"
	"os"
	"strings"
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
	MaxSize        int64 = 1024 * 1024 * 10 //10MB
)

func CloudStorageTriggerTest(ctx context.Context, e GCSEvent) error {
	if err := printMetaData(ctx, e); err != nil {
		return err
	}
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	bucket := client.Bucket(bucketName)

	tmpObjectName := fmt.Sprintf("/%s", e.Name)
	fmt.Printf("tmpObjectName: %v\n", tmpObjectName)
	dstObjectName := strings.Join(strings.Split(tmpObjectName, "/")[2:][:], "/")
	fmt.Printf("dstObjectName: %v\n", dstObjectName)
	dstObject := bucket.Object(dstObjectName)
	_, err = dstObject.Attrs(ctx)
	if err == nil {
		fmt.Printf("object: %s already exists", dstObjectName)
		return nil
	}
	if err != storage.ErrObjectNotExist {
		return err
	}

	srcObject := bucket.Object(e.Name)
	if err := validate(ctx, srcObject); err != nil {
		return err
	}
	if _, err := dstObject.CopierFrom(srcObject).Run(ctx); err != nil {
		return fmt.Errorf("dstObject.CopierFrom(srcObject).Run(ctx) - %v", err)
	}

	return nil
}

func validate(ctx context.Context, object *storage.ObjectHandle) error {
	objectAttrs, err := object.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("failed to get object attributes: %q, %v", object.ObjectName(), err)
	}

	if objectAttrs.Size > MaxSize {
		return fmt.Errorf("file is larger than %d, size: %d", MaxSize, objectAttrs.Size)
	}
	isSupportMIMEType := slice.Contains(mimeTypes, objectAttrs.ContentType)
	if isSupportMIMEType == false {
		return fmt.Errorf("unsupport MIME type, got = %q", objectAttrs.ContentType)
	}
	return nil
}

func printMetaData(ctx context.Context, e GCSEvent) error {
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
