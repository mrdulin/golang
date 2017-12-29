package services

type IAdPerformanceReportService interface {
	Get() string
}

type AdPerformanceReportService struct {
}

func NewAdPerformanceReportService() IAdPerformanceReportService {
	return &AdPerformanceReportService{}
}

func (svc *AdPerformanceReportService) Get() string {
	return "ad performance report"
}
