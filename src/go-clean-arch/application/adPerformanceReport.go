package application

import (
	"go-clean-arch/domain/services"
	"go-clean-arch/domain/services/adChannel/reports/adPerformance"
	"go-clean-arch/infrastructure/config"
	"log"
)

type IAdPerformanceReportUseCase interface {
	Get() error
}

type AdPerformanceReportUseCase struct {
	campaignService       services.ICampaignService
	campaignResultService services.ICampaignResultService
	appConfig             *config.ApplicationConfig
}

func NewAdPerformanceReportUseCase(
	campaignService services.ICampaignService,
	campaignResultService services.ICampaignResultService,
	appConfig *config.ApplicationConfig,
) IAdPerformanceReportUseCase {
	return &AdPerformanceReportUseCase{
		campaignService,
		campaignResultService,
		appConfig,
	}
}

func (uc *AdPerformanceReportUseCase) Get() error {
	googleCampaignIds, err := uc.campaignService.FindValidGoogleCampaignIds()
	if err != nil {
		return err
	}

	options := adPerformance.AdPerformanceReportServiceOptions{
		ClientCustomerId:      uc.appConfig.ClientCustomerId,
		RefreshToken:          uc.appConfig.RefreshToken,
		BaseURL:               uc.appConfig.AdChannelApi,
		CampaignResultService: uc.campaignResultService,
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
