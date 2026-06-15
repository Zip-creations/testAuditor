package jUnit


import "fmt"
import "os"
import "regexp"
import "encoding/xml"

func ReadGitNote(content string)(JUnitTestsuites, error) {
	data, _ := os.ReadFile(content)
	// This asssumes one file is provided, with a number of Junit XMl files mashed into one file by gitnotes, separated by an empty line
	// JUnit XML files CAN NOT include empty lines, since empty lines are used by gitnotes as default separator
	re := regexp.MustCompile(`\r?\n\r?\n`)
	parts := re.Split(string(data), -1)
	for _, part := range parts {
		fmt.Println("Part: ", part)
	}
	return ReadJUnitTestSuites(parts)
}

func ReadJUnitTestSuites(parts []string) (JUnitTestsuites, error) {
	var allSuites JUnitTestsuites
	for _, part := range parts {
		testSuites, err := ReadJUnitTestSuite(part)
		if err != nil {
			fmt.Println(err)  // TODO: log error somehow
			continue  // If one file is broken: skip and continue with the others
		}
		allSuites.Testsuites = append(allSuites.Testsuites, testSuites...)
	}
	return allSuites, nil
}

// func ReadJUnitTestSuites(path string) (JUnitTestsuites, error) {
// 	var allSuites JUnitTestsuites
// 	content, err := ReadDirRekursive(path)
// 	if err != nil {
// 		return allSuites, fmt.Errorf("Error reading directory:\n %s\n %w", path, err)
// 	}
// 	for _, filePath := range content {
// 		testSuites, err := ReadJUnitTestSuite(filePath)
// 		if err != nil {
// 			fmt.Println(err)  // TODO: log error somehow
// 			continue  // If one file is broken: skip and continue with the others
// 		}
// 		allSuites.Testsuites = append(allSuites.Testsuites, testSuites...)
// 	}
// 	return allSuites, nil
// }

// func ReadDirRekursive(path string) ([]string, error) {
// 	var allPaths []string
// 	entries, err := os.ReadDir(path)
// 	if err != nil {
// 		return allPaths, err
// 	}
// 	for _, entry := range entries {
// 		fullPath := filepath.Join(path, entry.Name())
// 		if entry.IsDir() {
// 			subPaths, err := ReadDirRekursive(fullPath)
// 			if err != nil {
// 				return allPaths, err
// 			}
// 			allPaths = append(allPaths, subPaths...)
// 		} else {
// 			allPaths = append(allPaths, fullPath)
// 		}
// 	}
// 	return allPaths, nil
// }

// func ReadJUnitTestSuite(filePath string) ([]JUnitTestsuite, error) {
// 	var testsuites JUnitTestsuites
// 	var testsuite JUnitTestsuite
// 	data, err := os.ReadFile(filePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("error while reading file:\n %s\n %w", filePath, err)
// 	}
// 	marshalErr1 := xml.Unmarshal(data, &testsuite)
// 	if marshalErr1 == nil {
// 		return []JUnitTestsuite{testsuite}, nil
// 	}
// 	marshalErr2 := xml.Unmarshal(data, &testsuites)
// 	if marshalErr2 == nil {
// 		return testsuites.Testsuites, nil
// 	}
// 	return testsuites.Testsuites, fmt.Errorf("Error while unmarshalling JUnit XML:\n %s\n %w\n %w", filePath, marshalErr1, marshalErr2)
// }

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
	return testsuites.Testsuites, fmt.Errorf("Error while unmarshalling JUnit XML:\n %s\n %w\n %w", marshalErr1, marshalErr2)
}
