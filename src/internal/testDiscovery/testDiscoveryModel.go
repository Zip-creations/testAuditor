package testDiscovery

import "encoding/xml"


func ParseTestDiscovery(data string) (DiscoveryTestsuite, error) {
	var suite DiscoveryTestsuite
	return suite, xml.Unmarshal([]byte(data), &suite)
}
