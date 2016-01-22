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

func AdPerformanceReport(_ context.Context, _ functions.PubSubMessage) error {
	usecase := application.NewAdPerformanceReportUseCase(compositionRoot.CampaignService)
	if err := usecase.Get(); err != nil {
		return err
	}
	return nil
}
