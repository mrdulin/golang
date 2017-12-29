package getCampaignReport

import (
	"cloud-function/03-multiple-functions-structure/domain/services"
	"fmt"
	"net/http"
)

func GetCampaignReport(w http.ResponseWriter, r *http.Request) {
	campaignService := services.NewCampaignReportService()
	report := campaignService.Get()
	if _, err := fmt.Fprint(w, report); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
