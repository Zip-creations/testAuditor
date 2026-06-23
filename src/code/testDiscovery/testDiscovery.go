package testDiscovery

import "fmt"
import "os/exec"
import "encoding/xml"
import cfg "github.com/Zip-creations/optimize_CI_deterministic_builds/src/code/config"

func RunTestDiscovery(command cfg.Command) (DiscoveryTestsuite, error) {
	out, err := exec.Command(command.Command, command.Args...).CombinedOutput()
	if err != nil {
		return DiscoveryTestsuite{}, fmt.Errorf("Error executing test discovery script: %w\n%s", err, out)
	}
	return XMLtoDiscoveryTestsuite([]byte(out), &DiscoveryTestsuite{})
}

func XMLtoDiscoveryTestsuite(data []byte, suite *DiscoveryTestsuite) (DiscoveryTestsuite, error) {
	err := xml.Unmarshal(data, suite)
	if err != nil {
		return *suite, fmt.Errorf("Error while unmarshalling user generated XML:\n %w", err)
	}
	return *suite, err
}
