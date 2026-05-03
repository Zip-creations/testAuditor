package main

import "encoding/xml"

type TestStatus string
const (
    StatusPassed  TestStatus = "passed"
    StatusFailed  TestStatus = "failed"
    StatusSkipped TestStatus = "skipped"
)

type Report struct {
	XMLName xml.Name `xml:"report"`
	Name string `xml:"name,attr"`
	TotalRun int `xml:"totalRun,attr"`
	TotalFailed int `xml:"totalFailed,attr"`
	TotalSkipped int `xml:"totalSkipped,attr"`
	TestSuites []TestSuiteReport `xml:"testSuiteReport"`
}

type TestSuiteReport struct {
	Name string `xml:"name,attr"`
	Timestamp string `xml:"timestamp,attr,omitempty"`
	TotalRunSuite int `xml:"totalRunSuite,attr"`
	TotalFailedSuite int `xml:"totalFailedSuite,attr"`
	TotalSkippedSuite int `xml:"totalSkippedSuite,attr"`	
	TestCases []TestCaseReport `xml:"testCaseReport"`
}

type TestCaseReport struct {
	Name string `xml:"name,attr"`
	Result TestStatus `xml:"result,attr"`
}
