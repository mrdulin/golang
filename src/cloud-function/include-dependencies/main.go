package includedependenciestest

import (
	"context"

	"cloud.google.com/go/storage"
)

func ListBudget() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {

	}
}
