package services

type ICampaignReportService interface {
	Get() string
}

type CampaignReportService struct {
}

func NewCampaignReportService() ICampaignReportService {
	return &CampaignReportService{}
}

func (svc *CampaignReportService) Get() string {
	return "campaign report"
}
