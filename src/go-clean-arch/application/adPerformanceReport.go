package application

import (
	"go-clean-arch/domain/services"
	"go-clean-arch/domain/services/adChannel/reports/adPerformance"
	"log"
	"os"
)

type IAdPerformanceReportUseCase interface {
	Get() error
}

type AdPerformanceReportUseCase struct {
	campaignService services.ICampaignService
}

func NewAdPerformanceReportUseCase(campaignService services.ICampaignService) IAdPerformanceReportUseCase {
	return &AdPerformanceReportUseCase{
		campaignService,
	}
}

func (uc *AdPerformanceReportUseCase) Get() error {
	googleCampaignIds, err := uc.campaignService.FindValidGoogleCampaignIds()
	if err != nil {
		return err
	}

	options := adPerformance.AdPerformanceReportServiceOptions{
		ClientCustomerId: 0,
		RefreshToken:     "",
		BaseURL:          os.Getenv("BaseURL"),
	}
	adPerformanceService := adPerformance.NewAdPerformanceReportService(options)
	reportDefinition := adPerformanceService.FormReportDefinition(googleCampaignIds)
	report, err := adPerformanceService.Get(reportDefinition)
	if err != nil {
		return err
	}
	reportRows := report.GetRows()
	err = adPerformanceService.UpdateStatusTransaction(reportRows)
	if err != nil {
		return err
	}

	log.Printf("update status transaction done")
	return nil
}
