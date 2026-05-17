package jUnit


import "fmt"
import "os"
import "encoding/xml"
import "path/filepath"


func ReadJUnitTestSuites(path string) (JUnitTestsuites, error) {
	var allSuites JUnitTestsuites
	content, err := ReadDirRekursive(path)
	if err != nil {
		return allSuites, fmt.Errorf("Error reading directory:\n %s\n %w", path, err)
	}
	for _, filePath := range content {
		testSuites, err := ReadJUnitTestSuite(filePath)
		if err != nil {
			fmt.Println(err)  // TODO: log error somehow
			continue  // If one file is broken: skip and continue with the others
		}
		allSuites.Testsuites = append(allSuites.Testsuites, testSuites...)
	}
	return allSuites, nil
}

func ReadDirRekursive(path string) ([]string, error) {
	var allPaths []string
	entries, err := os.ReadDir(path)
	if err != nil {
		return allPaths, err
	}
	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			subPaths, err := ReadDirRekursive(fullPath)
			if err != nil {
				return allPaths, err
			}
			allPaths = append(allPaths, subPaths...)
		} else {
			allPaths = append(allPaths, fullPath)
		}
	}
	return allPaths, nil
}

func ReadJUnitTestSuite(filePath string) ([]JUnitTestsuite, error) {
	var testsuites JUnitTestsuites
	var testsuite JUnitTestsuite
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error while reading file:\n %s\n %w", filePath, err)
	}
	marshalErr1 := xml.Unmarshal(data, &testsuite)
	if marshalErr1 == nil {
		return []JUnitTestsuite{testsuite}, nil
	}
	marshalErr2 := xml.Unmarshal(data, &testsuites)
	if marshalErr2 == nil {
		return testsuites.Testsuites, nil
	}
	return testsuites.Testsuites, fmt.Errorf("Error while unmarshalling JUnit XML:\n %s\n %w\n %w", filePath, marshalErr1, marshalErr2)
}
