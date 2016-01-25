package models

import "encoding/xml"

type ReportName struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type DateRange struct {
	Text string `xml:",chardata"`
	Date string `xml:"date,attr"`
}

type Table struct {
	Text string        `xml:",chardata"`
	Row  []interface{} `xml:"row"`
}

type BaseReport struct {
	XMLName    xml.Name   `xml:"report"`
	Text       string     `xml:",chardata"`
	ReportName ReportName `xml:"report-name"`
	DateRange  DateRange  `xml:"date-range"`
	Table      Table      `xml:"table"`
}
