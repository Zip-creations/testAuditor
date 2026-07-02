package parseinput

import "encoding/xml"


type TestAuditorInput struct {
	XMLName   xml.Name    `xml:"testAuditorInput"`
	TestDiscovery TestDiscovery		`xml:"testDiscovery"`
	Reports Reports	  `xml:"reports"`
}

type TestDiscovery struct {
	XMLName   xml.Name    `xml:"testDiscovery"`
	Content string `xml:",cdata"`
}

type Reports struct {
	XMLName   xml.Name    `xml:"reports"`
	Reports []Report `xml:"report"`
}

type Report struct {
	XMLName   xml.Name    `xml:"report"`
	Content string `xml:",cdata"`
}
