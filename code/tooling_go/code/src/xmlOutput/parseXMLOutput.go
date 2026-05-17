package xmlOutput

import shared "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/shared"
import "encoding/xml"


type Testsuites struct {
	XMLName   xml.Name    `xml:"testsuites"`
	Testsuites []Testsuite `xml:"testsuite"`
}

type Testsuite struct {
	Name      string     `xml:"name,attr"`
	Testcases []Testcase `xml:"testcase"`
}

type Testcase struct {
	Classname string `xml:"classname,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	QualifiedName string `xml:"qualifiedName,attr,omitempty"`
	Result TestStatus `xml:"result,attr"`  // if status is failed or skipped, additional info is added
	Failure *shared.Failure `xml:"failure,omitempty"`
	Skipped *shared.Skipped `xml:"skipped,omitempty"`
}

type TestStatus string
const (
    StatusPassed  TestStatus = "passed"
    StatusFailed  TestStatus = "failed"
    StatusSkipped TestStatus = "skipped"
	StatusNotExecuted TestStatus = "notExecuted"
)
