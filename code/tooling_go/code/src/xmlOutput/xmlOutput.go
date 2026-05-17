package xmlOutput

import disc "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/testDiscovery"
import junit "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/jUnit"


func MatchTests(discoverySuite disc.DiscoveryTestsuite, junitSuites junit.JUnitTestsuites) Testsuites {
	var matchedSuites Testsuites
	for _, testcaseXML := range discoverySuite.DiscoveryTestcases {
		testcase := Testcase{
			Classname: testcaseXML.Classname,
			Name: testcaseXML.Name,
			QualifiedName: testcaseXML.QualifiedName,
		}
		found := false
		for _, junitSuite := range junitSuites.Testsuites {
			for _, junitTestcase := range junitSuite.Testcases {
				if testcaseXML.Name == junitTestcase.Name && testcaseXML.Classname == junitTestcase.Classname {
					found = true
					testcase.Failure = junitTestcase.Failure
					testcase.Skipped = junitTestcase.Skipped
					if testcase.Skipped != nil {
						testcase.Result = StatusSkipped
					} else if testcase.Failure != nil {  // A test can't have been failed and skipped at the same time
						testcase.Result = StatusFailed
					} else {
						testcase.Result = StatusPassed
					}
					suit := FindTestsuiteByName(matchedSuites.Testsuites, junitSuite.Name)
					if suit == nil {
						matchedSuites.Testsuites = append(matchedSuites.Testsuites, Testsuite{Name: junitSuite.Name,})
						suit = &matchedSuites.Testsuites[len(matchedSuites.Testsuites)-1]
					}
					suit.Testcases = append(suit.Testcases, testcase)
					break
				}
			}
			if found {break}
		}
		if !found {
			testcase.Result = StatusNotExecuted
			// Group all tests that have not been executed
			neName := "not executed"
			suit := FindTestsuiteByName(matchedSuites.Testsuites, neName)
			if suit == nil {
				matchedSuites.Testsuites = append(matchedSuites.Testsuites, Testsuite{Name: neName,})
				suit = &matchedSuites.Testsuites[len(matchedSuites.Testsuites)-1]
			}
			suit.Testcases = append(suit.Testcases, testcase)
		}
	}
	return matchedSuites
}

func FindTestsuiteByName(suites []Testsuite, name string) *Testsuite {
	for i := range suites {
		if suites[i].Name == name {
			return &suites[i]
		}
	}
	return nil
}
