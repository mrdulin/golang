package services

import (
	models "go-clean-arch/domain/models/adChannel"
	"go-clean-arch/infrastructure/config"
	"log"
	"strconv"
	"testing"
)

var (
	options AdPerformanceReportServiceOptions
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
		BaseURL: viper.GetString("AD_CHANNEL_API"),
		ClientCustomerId: ClientCustomerId,
		RefreshToken: viper.GetString("RefreshToken"),
	}
	adPerformanceService = NewAdPerformanceReportService(options)
}

func TestGet(t *testing.T) {
	t.Skip("skip TestGet")
	report, err := adPerformanceService.Get()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(report)
}

func TestGetWithReportDefinition(t *testing.T) {
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