package testDiscovery


type DiscoveryTestsuite struct {
	DiscoveryTestcases []DiscoveryTestcase `xml:"testcase"`
}

type DiscoveryTestcase struct {
	Classname string `xml:"classname,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	QualifiedName string `xml:"qualifiedName,attr,omitempty"`
}
