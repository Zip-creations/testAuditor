package jUnit

import shared "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/shared"
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
	Failure *shared.Failure `xml:"failure,omitempty"`
	Skipped *shared.Skipped `xml:"skipped,omitempty"`
}
