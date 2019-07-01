package reports

import (
	"context"
	"go-clean-arch/domain/models/gcloud/functions"
	"testing"
)

func TestAdPerformanceReport(t *testing.T) {
	ctx := context.Background()
	m := functions.PubSubMessage{}
	err := GetAdPerformanceReport(ctx, m)
	if err != nil {
		t.Logf("%+v", err)
	}
}
