package xmlOutput

import disc "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/testDiscovery"
import rep "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/report"


type testKey struct {
	Classname string
	Name      string
}

func MatchTests(discoverySuite disc.DiscoveryTestsuite, junitSuites rep.JUnitTestsuites) []string {
	seen := make(map[testKey]struct{})

	for _, junitSuite := range junitSuites.Testsuites {
		for _, junitTestcase := range junitSuite.Testcases {
			seen[testKey{
				Classname: junitTestcase.Classname,
				Name:      junitTestcase.Name,
			}] = struct{}{}
		}
	}

	result := make([]string, 0, len(discoverySuite.DiscoveryTestcases))

	for _, testcaseXML := range discoverySuite.DiscoveryTestcases {
		key := testKey{
			Classname: testcaseXML.Classname,
			Name:      testcaseXML.Name,
		}

		if _, found := seen[key]; !found {
			result = append(result, testcaseXML.QualifiedName)
		}
	}
	return result
}
