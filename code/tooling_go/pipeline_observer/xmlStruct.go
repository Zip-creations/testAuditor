package main

// TODO: The XML structure can have two different root elements: <testsuites> and <testsuite>. Tool should be able to process both.
// Currently, <testsuite> is assumed as root
type Testsuites struct {
	Testsuites []Testsuite `xml:"testsuites"`
}

type Testsuite struct {
	Name      string     `xml:"name,attr"`
	Testcases []Testcase `xml:"testcase"`
}

type Testcase struct {
	Classname string `xml:"classname,attr"`
	Name string `xml:"name,attr"`
	Failure *Failure `xml:"failure,omitempty"`
	Skipped *bool `xml:"skipped,omitempty"`
}

type Failure struct {
	Message string `xml:"message,attr"`
	Type string `xml:"type,attr"`
	Content string `xml:",chardata"`
}

func (t Testcase) IsSkipped() bool {
    return t.Skipped != nil
}
