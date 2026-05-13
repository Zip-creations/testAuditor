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
	TestsTotal int `xml:"testsTotal,attr"`
	TestsRun int `xml:"testsRun,attr"`
	// TotalFailed int `xml:"totalFailed,attr"`
	TestsSkipped int `xml:"testsSkipped,attr"`
	TestSuites []TestSuiteReport `xml:"testSuiteReport"`
}

type TestSuiteReport struct {
	Name string `xml:"name,attr"`
	Timestamp string `xml:"timestamp,attr,omitempty"`
	TestsTotal int `xml:"testsTotal,attr"`
	TestsRun int `xml:"testsRun,attr"`
	// TestsFailedSuite int `xml:"testsFailedSuite,attr"`
	TestsSkipped int `xml:"testsSkipped,attr"`
	TestCases []TestCaseReport `xml:"testCaseReport"`
}

type TestCaseReport struct {
	Name string `xml:"name,attr"`
	Result TestStatus `xml:"result,attr"`
}
