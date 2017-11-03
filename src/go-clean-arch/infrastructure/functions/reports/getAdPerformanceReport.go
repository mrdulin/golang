package reports

import (
	"context"
	"go-clean-arch/application"
	"go-clean-arch/domain/models/googleCloud/functions"
)

var compositionRoot *application.CompositionRoot

func init() {
	compositionRoot = application.NewCompositionRoot()
}

func GetAdPerformanceReport(_ context.Context, _ functions.PubSubMessage) error {
	adPerformanceReportUseCase := application.NewAdPerformanceReportUseCase(compositionRoot.CampaignService, compositionRoot.CampaignResultService, compositionRoot.AppConfig)
	if err := adPerformanceReportUseCase.Get(); err != nil {
		return err
	}
	return nil
}
