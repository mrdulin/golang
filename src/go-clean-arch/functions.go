package serverlessfns

import (
	"context"
	"go-clean-arch/domain/models/gcloud/functions"
	"go-clean-arch/infrastructure/functions/reports"
)

func GetAdPerformanceReport(ctx context.Context, m functions.PubSubMessage) error {
	return reports.GetAdPerformanceReport(ctx, m)
}
