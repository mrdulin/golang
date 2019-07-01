package cloudrun

import (
	"go-clean-arch/application"
	"net/http"
)

var (
	compositionRoot            *application.CompositionRoot
	adPerformanceReportUseCase application.IAdPerformanceReportUseCase
)

func init() {
	compositionRoot = application.NewCompositionRoot()
	adPerformanceReportUseCase = application.NewAdPerformanceReportUseCase(
		compositionRoot.CampaignService,
		compositionRoot.CampaignResultService,
		compositionRoot.AppConfig,
	)
}

func GetAdPerformanceReport(w http.ResponseWriter, r *http.Request) {
	if err := adPerformanceReportUseCase.Get(); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
