package adPerformance

import (
	models "go-clean-arch/domain/models/adChannel"
	"go-clean-arch/infrastructure/config"
	"log"
	"strconv"
	"testing"
)

var (
	options              AdPerformanceReportServiceOptions
	adPerformanceService IAdPerformanceReportService
)

func init() {
	appConfig := config.NewApplicationConfig()
	viper, err := appConfig.Load()
	if err != nil {
		log.Fatal(err)
	}
	ClientCustomerId, err := strconv.Atoi(viper.GetString("ClientCustomerId"))
	if err != nil {
		log.Fatal(err)
	}
	options = AdPerformanceReportServiceOptions{
		BaseURL:          viper.GetString("AD_CHANNEL_API"),
		ClientCustomerId: ClientCustomerId,
		RefreshToken:     viper.GetString("RefreshToken"),
	}
	adPerformanceService = NewAdPerformanceReportService(options)
}

func TestGet(t *testing.T) {
	t.Skip()
	report, err := adPerformanceService.Get()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(report)
}

func TestGetWithFields(t *testing.T) {
	t.Skip()
	reportDefinition := models.ReportDefinition{
		Selector: models.Selector{
			Fields: []string{"CampaignId"},
		},
	}
	report, err := adPerformanceService.Get(reportDefinition)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(report)
}

func TestGetWithPredicate(t *testing.T) {
	t.Skip()
	campaignIds := []int{1898036877, 1900752711, 2029744960, 1900528765, 2039653855, 1897998186, 1900833347, 1915846792}
	reportDefinition := adPerformanceService.FormReportDefinition(campaignIds)
	report, err := adPerformanceService.Get(reportDefinition)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", report)
}

func TestGetRowsOfReport(t *testing.T) {
	campaignIds := []int{1898036877, 1900752711, 2029744960, 1900528765, 2039653855, 1897998186, 1900833347, 1915846792}
	reportDefinition := adPerformanceService.FormReportDefinition(campaignIds)
	report, err := adPerformanceService.Get(reportDefinition)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	rows := report.GetRows()
	t.Logf("%+v", rows)
}
