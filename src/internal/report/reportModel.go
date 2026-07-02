package jUnit

import "encoding/xml"


type JUnitTestsuites struct {
	XMLName   xml.Name    `xml:"testsuites"`
	Testsuites []JUnitTestsuite `xml:"testsuite"`
}

type JUnitTestsuite struct {
	XMLName   xml.Name    `xml:"testsuite"`
	Name      string     `xml:"name,attr"`
	Testcases []JUnitTestcase `xml:"testcase"`
}

type JUnitTestcase struct {
	Classname string `xml:"classname,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	Failure *Failure `xml:"failure,omitempty"`
	Skipped *Skipped `xml:"skipped,omitempty"`
}

type Failure struct {
	Message string `xml:"message,attr"`
	Type string `xml:"type,attr"`
	Content string `xml:",chardata"`
}

type Skipped struct {
	Message string `xml:"message,attr"`
}
