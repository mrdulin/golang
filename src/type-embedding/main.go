package main

import (
  "fmt"
  "type-embedding/report"
)

func main() {
  campaignPerformanceReport := report.CampaignPerformanceReport{}
  reportService := report.Service{BaseURL: "https://github.com", RefreshToken: "123"}
  content, err := (&reportService).Get("reportDefinition", campaignPerformanceReport)
  if err != nil {
    fmt.Printf("err=%v", err)
  }
  fmt.Printf("content=%v", content)
}
