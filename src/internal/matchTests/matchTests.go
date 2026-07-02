package xmlOutput

import disc "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/testDiscovery"
import rep "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/report"


func MatchTests(discoverySuite disc.DiscoveryTestsuite, junitSuites rep.JUnitTestsuites) []string {
	var result []string
	for _, testcaseXML := range discoverySuite.DiscoveryTestcases {
		found := false
		for _, junitSuite := range junitSuites.Testsuites {
			for _, junitTestcase := range junitSuite.Testcases {
				if testcaseXML.Name == junitTestcase.Name && testcaseXML.Classname == junitTestcase.Classname {
					found = true
					break
				}
			}
			if found {break}
		}
		if !found {
			result = append(result, testcaseXML.QualifiedName)
		}
	}
	return result
}
