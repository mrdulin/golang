package cloudrun

import (
	"encoding/json"
	"go-clean-arch/application"
	"go-clean-arch/domain/models/gcloud/cloudrun"
	"log"
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
	var m cloudrun.PubSubMessage

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Printf("json.NewDecoder: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := adPerformanceReportUseCase.Get(); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
