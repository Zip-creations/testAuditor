package xmlOutput

import disc "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/testDiscovery"
import rep "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/report"


func MatchTests(discoverySuite disc.DiscoveryTestsuite, seen map[rep.TestKey]struct{}) []string {
	result := make([]string, 0, len(discoverySuite.DiscoveryTestcases))
	for _, testcaseXML := range discoverySuite.DiscoveryTestcases {
		key := rep.TestKey{
			Classname: testcaseXML.Classname,
			Name:      testcaseXML.Name,
		}
		if _, found := seen[key]; !found {
			result = append(result, testcaseXML.QualifiedName)
		}
	}
	return result
}
