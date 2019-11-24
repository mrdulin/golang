package main

import (
  "bytes"
  "fmt"
  "html/template"
)

func main() {
  query := `
    INSERT INTO "ADGROUP_PERFORMANCE_REPORT" (
      {{.columnPrefix}}_adgroup_id, 
      {{.columnPrefix}}_adgroup_nme, 
      {{.columnPrefix}}_adgroup_status, 
      {{.columnPrefix}}_campaign_id, 
      {{.columnPrefix}}_campaign_nme, 
      {{.columnPrefix}}_campaign_status, 
      {{.columnPrefix}}_clicks, 
      {{.columnPrefix}}_impressions, 
      {{.columnPrefix}}_ctr, 
      {{.columnPrefix}}_average_cpc, 
      {{.columnPrefix}}_cost, 
      {{.columnPrefix}}_conversions, 
      {{.columnPrefix}}_average_position, 
      {{.columnPrefix}}_device, 
      google_adwords_client_customer_id
    ) VALUES
  `
  vars := make(map[string]interface{})
  vars["columnPrefix"] = "adgroup_performance_report"
  result := processString(query, vars)
  fmt.Printf("result=%s\n", result)
}

func process(t *template.Template, vars interface{}) string {
  var tmplBytes bytes.Buffer
  err := t.Execute(&tmplBytes, vars)
  if err != nil {
    panic(err)
  }
  return tmplBytes.String()
}

func processString(str string, vars interface{}) string {
  tmpl, err := template.New("tmpl").Parse(str)
  if err != nil {
    panic(err)
  }
  return process(tmpl, vars)
}
