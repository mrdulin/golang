package adPerformance

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	models "go-clean-arch/domain/models/adChannel"
	"go-clean-arch/domain/services"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type IAdPerformanceReportService interface {
	Get(reportDefinition ...models.ReportDefinition) (*models.AdPerformanceReport, error)
	FormReportDefinition(campaignIds []int) models.ReportDefinition
	UpdateStatusTransaction(rows []models.AdPerformanceReportRow) error
}

type AdPerformanceReportServiceOptions struct {
	ClientCustomerId int
	RefreshToken     string
	BaseURL          string

	CampaignResultService services.ICampaignResultService
}

type AdPerformanceReportService struct {
	options AdPerformanceReportServiceOptions
}

func NewAdPerformanceReportService(options AdPerformanceReportServiceOptions) IAdPerformanceReportService {
	options.BaseURL = options.BaseURL + "/google/report/AdPerformance"
	return &AdPerformanceReportService{options}
}

func (svc *AdPerformanceReportService) Get(reportDefinition ...models.ReportDefinition) (*models.AdPerformanceReport, error) {
	var body = models.ReportDefinition{
		DateRangeType: "ALL",
	}

	if len(reportDefinition) != 0 {
		body = reportDefinition[0]
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, errors.Wrap(err, "json.NewEncoder(buf).Encode(body)")
	}

	req, err := http.NewRequest("POST", svc.options.BaseURL, buf)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest")
	}
	req.Header.Set("ClientCustomerId", strconv.Itoa(svc.options.ClientCustomerId))
	req.Header.Set("RefreshToken", svc.options.RefreshToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client.Do(req)")
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll(resp.Body)")
	}

	report := models.AdPerformanceReport{}
	// TODO: content = {code: 2000, message: 'xxx'}
	err = xml.Unmarshal(content, &report)
	if err != nil {
		return nil, errors.Wrap(err, "xml.Unmarshal(content, &report)")
	}

	return &report, nil
}

func (svc *AdPerformanceReportService) FormReportDefinition(campaignIds []int) models.ReportDefinition {
	var campaignStrIds []string

	for _, campaignId := range campaignIds {
		campaignStrIds = append(campaignStrIds, strconv.Itoa(campaignId))
	}

	return models.ReportDefinition{
		Selector: models.Selector{
			Fields: []string{
				"AdGroupName",
				"AdGroupId",
				"AdGroupStatus",
				"AdType",
				"CampaignId",
				"CampaignName",
				"CampaignStatus",
				"CombinedApprovalStatus",
				"Id",
				"Status",
			},
			Predicates: []models.Predicate{
				{
					Field:    "CampaignId",
					Operator: models.IN, // TODO: add namespace like: models.Operator.IN
					Values:   campaignStrIds,
				},
			},
		},
	}
}

// UpdateStatusTransaction will update campaign status and ad group status, ad status.
func (svc *AdPerformanceReportService) UpdateStatusTransaction(rows []models.AdPerformanceReportRow) error {
	err := svc.options.CampaignResultService.UpdateStatusTransaction(rows)
	if err != nil {
		return errors.Wrap(err, "svc.options.campaignResultRepo.UpdateStatusTransaction()")
	}
	return nil
}
