package jUnit

import "fmt"
import "encoding/xml"
import input "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/input"


func ParseJUnitTestSuites(reports input.Reports) (JUnitTestsuites, error) {
	var allSuites JUnitTestsuites
	for _, part := range reports.Reports {
		testSuites, err := ParseJUnitTestSuite(part.Content)
		if err != nil {
			fmt.Println(err)  // TODO: log error somehow
			continue  // If one section is broken: skip and continue with the others
		}
		allSuites.Testsuites = append(allSuites.Testsuites, testSuites...)
	}
	return allSuites, nil
}

func ParseJUnitTestSuite(part string) ([]JUnitTestsuite, error) {
	var testsuites JUnitTestsuites
	var testsuite JUnitTestsuite
	// Since both tags can appear as root tag in JUnit-XMl, try both if the first couldn't be parsed
	marshalErr1 := xml.Unmarshal([]byte(part), &testsuite)
	if marshalErr1 == nil {
		return []JUnitTestsuite{testsuite}, nil
	}
	marshalErr2 := xml.Unmarshal([]byte(part), &testsuites)
	if marshalErr2 == nil {
		return testsuites.Testsuites, nil
	}
	return testsuites.Testsuites, fmt.Errorf("Error while unmarshalling JUnit XML:\n %w\n %w", marshalErr1, marshalErr2)
}
