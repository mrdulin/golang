package repositories

import models "go-clean-arch/domain/models/adChannel"

type CampaignResultRepository interface {
	UpdateStatusTransaction(row models.AdPerformanceReportRow) error
}
