package getAdPerformanceReport

import (
	"cloud-function/03-multiple-functions-structure/domain/services"
	"fmt"
	"net/http"
)

func GetAdPerformanceReport(w http.ResponseWriter, r *http.Request) {
	adPerformanceReport := services.NewAdPerformanceReportService()
	report := adPerformanceReport.Get()
	if _, err := fmt.Fprint(w, report); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
