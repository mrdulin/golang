package services

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/pkg/errors"
	"go-clean-arch/domain/models/adChannel"
	"io/ioutil"
	"net/http"
	"strconv"
)

type IAdPerformanceReportService interface {
	Get(reportDefinition ...models.ReportDefinition) (*models.AdPerformanceReport, error)
}

type AdPerformanceReportServiceOptions struct {
	ClientCustomerId int
	RefreshToken string
	BaseURL string
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
	err = xml.Unmarshal(content, &report)
	if err != nil {
		return nil, errors.Wrap(err, "xml.Unmarshal(content, &report)")
	}

	return &report, nil
}