package jUnit


import "fmt"
import "regexp"
import "encoding/xml"

func ReadGitNote(content string)(JUnitTestsuites, error) {
	// This asssumes one file is provided, with a number of Junit XMl files mashed into one file by gitnotes, separated by an empty line
	// JUnit XML files CAN NOT include empty lines, since empty lines are used by gitnotes as default separator!
	if content == "" {
		return JUnitTestsuites{
			XMLName: xml.Name{Local: "default"},
			Testsuites: []JUnitTestsuite{},
		}, nil
	}
	re := regexp.MustCompile(`\r?\n\r?\n`)
	parts := re.Split(content, -1)
	// for _, part := range parts {
	// 	fmt.Println("Part: ", part)
	// }
	return ReadJUnitTestSuites(parts)
}

func ReadJUnitTestSuites(parts []string) (JUnitTestsuites, error) {
	var allSuites JUnitTestsuites
	for _, part := range parts {
		testSuites, err := ReadJUnitTestSuite(part)
		if err != nil {
			fmt.Println(err)  // TODO: log error somehow
			continue  // If one section is broken: skip and continue with the others
		}
		allSuites.Testsuites = append(allSuites.Testsuites, testSuites...)
	}
	return allSuites, nil
}

func ReadJUnitTestSuite(part string) ([]JUnitTestsuite, error) {
	var testsuites JUnitTestsuites
	var testsuite JUnitTestsuite
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
