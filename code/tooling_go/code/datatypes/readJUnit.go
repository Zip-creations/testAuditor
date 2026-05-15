package datatypes


type JUnitTestsuites struct {
	Testsuites []JUnitTestsuite `xml:"testsuites"`
}

type JUnitTestsuite struct {
	Name      string     `xml:"name,attr"`
	Timestamp string     `xml:"timestamp,attr,omitempty"`
	Testcases []JUnitTestcase `xml:"testcase"`
}

type JUnitTestcase struct {
	Classname string `xml:"classname,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	Failure *Failure `xml:"failure,omitempty"`
	Skipped *Skipped `xml:"skipped,omitempty"`
}
