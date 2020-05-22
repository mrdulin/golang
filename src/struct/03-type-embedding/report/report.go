package report

import "fmt"

type Reporter interface {
	Get(reportDefinition interface{}, model interface{}) (interface{}, error)
}

type Service struct {
	BaseURL      string
	RefreshToken string
}

func (svc *Service) Get(reportDefinition interface{}, model interface{}) (interface{}, error) {
	fmt.Printf("BaseURL=%v\n", svc.BaseURL)
	fmt.Printf("RefreshToken=%v\n", svc.RefreshToken)
	return model, nil
}
