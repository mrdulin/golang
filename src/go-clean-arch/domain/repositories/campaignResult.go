package repositories

import (
	adChannelModels "go-clean-arch/domain/models/adChannel"
	"go-clean-arch/domain/models/cedar/campaign"
)

type CampaignResultRepository interface {
	UpdateStatusTransaction(row adChannelModels.AdPerformanceReportRow, status campaign.CampaignChannelStatus) error
}
