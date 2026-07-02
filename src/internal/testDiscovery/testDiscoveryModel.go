package testDiscovery

import "fmt"
import "encoding/xml"

func XMLtoDiscoveryTestsuite(data string) (DiscoveryTestsuite, error) {
	var suite DiscoveryTestsuite
	err := xml.Unmarshal([]byte(data), &suite)
	if err != nil {
		return suite, fmt.Errorf("Error while unmarshalling user generated XML:\n %w", err)
	}
	return suite, err
}
