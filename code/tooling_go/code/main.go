package main

import "fmt"
import "os"
import "os/exec"
import "encoding/xml"
import "encoding/json"
import "path/filepath"
import dt "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/pipeline_observer/datatypes"

func main() {
	// Read config
	config, err := ReadConfig("./config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config, "\n")  // Debug
	
	// Read all existing tests from the user-configured script
	allSuites, err := RunTestDiscoveryScript(config.TestDiscoveryPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(allSuites, "\n")  // Debug

	// Read all tests in the JUnit XML output of the last run (if existing)
	allSuitesJUnit, err := ReadJUnitTestSuites(config.JUnitXMLDirectory)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(allSuitesJUnit, "\n")  // Debug

	report := MatchTests(allSuites, allSuitesJUnit)
	WriteXMLToFile(report, config.OutputPath)
	fmt.Println("Successfully created report: \n", report)  // Debug
}

func ReadConfig(path string) (dt.Config, error) {
	var config dt.Config
	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("Error while reading config file:\n %s\n %w", path, err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("Error while unmarshalling config JSON:\n %w", err)
	}
	return config, nil
}

func RunTestDiscoveryScript(path string) (dt.DiscoveryTestsuite, error) {
	cmd := exec.Command("bash", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return dt.DiscoveryTestsuite{}, fmt.Errorf("Error executing test discovery script: %w\n%s", err, out)
	}
	return XMLtoDiscoveryTestsuite([]byte(out), &dt.DiscoveryTestsuite{})
}

func XMLtoDiscoveryTestsuite(data []byte, suite *dt.DiscoveryTestsuite) (dt.DiscoveryTestsuite, error) {
	err := xml.Unmarshal(data, suite)
	if err != nil {
		return *suite, fmt.Errorf("Error while unmarshalling user generated XML:\n %w", err)
	}
	return *suite, err
}

func ReadJUnitTestSuites(path string) (dt.JUnitTestsuites, error) {
	var allSuites dt.JUnitTestsuites
	content, err := os.ReadDir(path)
	if err != nil {
		return allSuites, fmt.Errorf("Error reading directory:\n %s\n %w", path, err)
	}
	content = filterForXML(content)
	for _, entry := range content {
		filePath := filepath.Join(path, entry.Name())
		testSuit, err := ReadJUnitTestSuite(filePath)
		if err != nil {
			fmt.Println(err)  // TODO: log error somehow
			continue  // If one file is broken: skip and continue with the others
		}
		allSuites.Testsuites = append(allSuites.Testsuites, testSuit)
	}
	return allSuites, nil
}

func ReadJUnitTestSuite(filePath string) (dt.JUnitTestsuite, error) {
	var testsuite dt.JUnitTestsuite
	data, err := os.ReadFile(filePath)
	if err != nil {
		return testsuite, fmt.Errorf("Error while reading file:\n %s\n %w", filePath, err)
	}
	return XMLtoJUnitTestSuite(data, &testsuite)
}

func XMLtoJUnitTestSuite(data []byte, suite *dt.JUnitTestsuite) (dt.JUnitTestsuite, error) {
	err := xml.Unmarshal(data, suite)
	if err != nil {
		return *suite, fmt.Errorf("Error while unmarshalling TestSuite XML:\n %w", err)
	}
	return *suite, err
}

func filterForXML(files []os.DirEntry) []os.DirEntry {
	var xmlFiles []os.DirEntry
	for _, file := range files {
		if file.Name()[len(file.Name())-4:] == ".xml" {
			xmlFiles = append(xmlFiles, file)
		}
	}
	return xmlFiles
}

func MatchTests(discoverySuite dt.DiscoveryTestsuite, junitSuites dt.JUnitTestsuites) dt.Testsuites {
	var matchedSuites dt.Testsuites
	for _, testcaseXML := range discoverySuite.DiscoveryTestcases {
		testcase := dt.Testcase{
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
						testcase.Result = dt.StatusSkipped
					} else if testcase.Failure != nil {  // A test can't have been failed and skipped at the same time
						testcase.Result = dt.StatusFailed
					} else {
						testcase.Result = dt.StatusPassed
					}
					suit := FindTestsuiteByName(matchedSuites.Testsuites, junitSuite.Name)
					if suit == nil {
						matchedSuites.Testsuites = append(matchedSuites.Testsuites, dt.Testsuite{Name: junitSuite.Name,})
						suit = &matchedSuites.Testsuites[len(matchedSuites.Testsuites)-1]
					}
					suit.Testcases = append(suit.Testcases, testcase)
					break
				}
			}
			if found {break}
		}
		if !found {
			testcase.Result = dt.StatusNotExecuted
			// Group all tests that have not been executed
			suit := FindTestsuiteByName(matchedSuites.Testsuites, "Not executed")
			if suit == nil {
				matchedSuites.Testsuites = append(matchedSuites.Testsuites, dt.Testsuite{Name: "Not executed",})
				suit = &matchedSuites.Testsuites[len(matchedSuites.Testsuites)-1]
			}
			suit.Testcases = append(suit.Testcases, testcase)
		}
	}
	return matchedSuites
}

func FindTestsuiteByName(suites []dt.Testsuite, name string) *dt.Testsuite {
	for i := range suites {
		if suites[i].Name == name {
			return &suites[i]
		}
	}
	return nil
}

func WriteXMLToFile(report dt.Testsuites, filePath string) error {
	data, err := xml.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("Error while marshalling Report XML:\n %w", err)
	}
	return os.WriteFile(filePath, data, 0644)
}
