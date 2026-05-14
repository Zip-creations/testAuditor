package main

import "fmt"
import "os"
import "os/exec"
import "encoding/xml"
import "path/filepath"
import "bytes"
import dt "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/pipeline_observer/datatypes"

func main() {
	// Read all existing tests from the user-configured script
	output, err := RunTestDiscoveryScript("examples/sample_find.sh")  // TODO: Read from config.json
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output, "\n")  // Debug

	// Read all tests in the JUnit XML output of the last run (if existing)
	allSuites, err := ReadJUnitTestSuites("./examples/jUnit_XML")  // TODO: Read from config.json
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(allSuites, "\n")  // Debug

	// report := CreateReport("Test Report", allSuites)
	// WriteXMLToFile(report, "./out/report.xml")
	// fmt.Println("Successfully created report: \n", report)  // Debug
}

func RunTestDiscoveryScript(path string) (dt.DiscoveryTestsuite, error) {
	cmd := exec.Command("bash", path)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return dt.DiscoveryTestsuite{}, fmt.Errorf("Error executing test discovery script: %w\n%s", err, stderr.String())
	}
	return XMLtoDiscoveryTestsuite([]byte(out.String()), &dt.DiscoveryTestsuite{})
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

// ~~~~~~~~~

func WriteXMLToFile(report Report, filePath string) error {
	data, err := xml.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("Error while marshalling Report XML:\n %w", err)
	}
	return os.WriteFile(filePath, data, 0644)
}
