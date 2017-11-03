package services

import (
	"fmt"
	adChannelModels "go-clean-arch/domain/models/adChannel"

	googleChannelAd "go-clean-arch/domain/models/adChannel/ad"
	googleChannelAdGroup "go-clean-arch/domain/models/adChannel/adGroup"
	googleChannelCampaign "go-clean-arch/domain/models/adChannel/campaign"
	cedarCampaign "go-clean-arch/domain/models/cedar/campaign"

	"go-clean-arch/domain/repositories"
)

type ICampaignResultService interface {
	UpdateStatusTransaction(rows []adChannelModels.AdPerformanceReportRow) error
}

type CampaignResultService struct {
	campaignResultRepo repositories.CampaignResultRepository
}

func NewCampaignResultService(campaignResultRepo repositories.CampaignResultRepository) ICampaignResultService {
	return &CampaignResultService{campaignResultRepo}
}

func (svc *CampaignResultService) UpdateStatusTransaction(rows []adChannelModels.AdPerformanceReportRow) error {
	// TODO: batch update
	for _, row := range rows {
		var campaignChannelStatus cedarCampaign.CampaignChannelStatus = ""

		if row.ApprovalStatus == googleChannelAd.APPROVED {

			if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.ENABLED && row.AdState == googleChannelAd.ENABLED {
				campaignChannelStatus = cedarCampaign.SUCCESS
			} else if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.PAUSED && row.AdState == googleChannelAd.ENABLED {
				campaignChannelStatus = cedarCampaign.FAILED
			} else if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.ENABLED && row.AdState == googleChannelAd.PAUSED {
				campaignChannelStatus = cedarCampaign.FAILED
			} else if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.REMOVED && row.AdState == googleChannelAd.ENABLED {
				campaignChannelStatus = cedarCampaign.FAILED
				// TODO: ad doesn't have "removed" status
			} else if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.ENABLED && row.AdState == googleChannelAd.DISABLED {
				campaignChannelStatus = cedarCampaign.FAILED
			} else if row.CampaignState == googleChannelCampaign.PAUSED && row.AdGroupState == googleChannelAdGroup.ENABLED && row.AdState == googleChannelAd.ENABLED {
				campaignChannelStatus = cedarCampaign.SUCCESS
			} else if row.CampaignState == googleChannelCampaign.REMOVED {
				campaignChannelStatus = cedarCampaign.FAILED
			}
			// TODO: campaign ended

		} else if row.ApprovalStatus == googleChannelAd.DISAPPROVED {
			if row.CampaignState == googleChannelCampaign.ENABLED && row.AdGroupState == googleChannelAdGroup.ENABLED && row.AdState == googleChannelAd.ENABLED {
				campaignChannelStatus = cedarCampaign.FAILED
			}
		}

		if campaignChannelStatus == "" {
			return fmt.Errorf("invalid campaign channel status")
		}

		return svc.campaignResultRepo.UpdateStatusTransaction(row, campaignChannelStatus)
	}
	return nil
}
