package services

import (
	"go-clean-arch/infrastructure/config"
	"log"
	"strconv"
	"testing"
)

//var options = AdPerformanceReportServiceOptions{
//	BaseURL: "http://localhost:3200/ad/api/v1",
//	ClientCustomerId: 9258066191,
//	RefreshToken: "1/P0z-iBX65sj-O666TnftiE9VSvb927m39SKG9D3Y_jI",
//}

var options AdPerformanceReportServiceOptions

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
}

func TestGet(t *testing.T) {
	adPerformanceService := NewAdPerformanceReportService(options)
	report, err := adPerformanceService.Get()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(report)
}