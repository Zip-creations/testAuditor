package main

import "fmt"
import "os"
import "os/exec"
import "encoding/xml"
import "path/filepath"
import "bytes"

func main() {
	allSuites, err := ReadTestSuites("./examples/jUnit_XML")  // TODO: ask path on first start of tool. NiceToHave: Make it configurable
	if err != nil {
		fmt.Println(err)
		return
	}
	output, err := RunTestDiscoveryScript("examples/sample_find.sh")
	fmt.Println(output, "\n")  // Debug
	report := CreateReport("Test Report", allSuites)
	WriteXMLToFile(report, "./out/report.xml")
	fmt.Println("Successfully created report: \n", report)  // Debug
}

func RunTestDiscoveryScript(path string) (Testsuite, error) {
	cmd := exec.Command("bash", path)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return Testsuite{}, fmt.Errorf("Error executing test discovery script: %w\n%s", err, stderr.String())
	}
	return XMLtoTestSuite([]byte(out.String()), &Testsuite{})
}

func ReadTestSuites(path string) (Testsuites, error) {
	var allSuites Testsuites
	content, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading files: ", err)
		return allSuites, err
	}
	content = filterForXML(content)
	for _, entry := range content {
		filePath := filepath.Join(path, entry.Name())
		testSuit, err := ReadTestSuite(filePath)
		if err != nil {
			fmt.Println(err)  // TODO: log error somehow
			continue  // If one file is broken: skip and continue with the others
		}
		allSuites.Testsuites = append(allSuites.Testsuites, testSuit)
	}
	return allSuites, nil
}

func ReadTestSuite(filePath string) (Testsuite, error) {
	var testsuite Testsuite
	data, err := os.ReadFile(filePath)
	if err != nil {
		return testsuite, fmt.Errorf("Error while reading file:\n %s\n %w", filePath, err)
	}
	return XMLtoTestSuite(data, &testsuite)
}

func XMLtoTestSuite(data []byte, suite *Testsuite) (Testsuite, error) {
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

func CreateReport(name string, testSuites Testsuites) Report {
	var totalRun, totalFailed, totalSkipped int = 0, 0, 0
	allSuites := []TestSuiteReport{}
	for _, testsuite := range testSuites.Testsuites {
		allSuites = append(allSuites, CreateTestSuiteReport(testsuite, &totalRun, &totalFailed, &totalSkipped))
	}
	return Report{
		Name: name,
		TestsTotal: totalRun + totalSkipped,
		TestsRun: totalRun,
		// TestsFailed: totalFailed,
		TestsSkipped: totalSkipped,
		TestSuites: allSuites,
	}
}

func CreateTestSuiteReport(testsuite Testsuite, totalRun *int, totalFailed *int, totalSkipped *int) TestSuiteReport {
	var testCases []TestCaseReport
	var totalRunSuite, totalFailedSuite, totalSkippedSuite int = 0, 0, 0
	for _, testcase := range testsuite.Testcases {
		var result TestStatus
		if testcase.IsSkipped() {
			*totalSkipped++
			totalSkippedSuite++
			result = StatusSkipped
		} else if testcase.HasFailed() {
			*totalFailed++
			totalFailedSuite++
			result = StatusFailed
		} else {
			result = StatusPassed
		}
		*totalRun++
		totalRunSuite++
		testCases = append(testCases, TestCaseReport{
			Name:   testcase.Name,
			Result: result,
		})
	}
	return TestSuiteReport{
		Name:          testsuite.Name,
		Timestamp: testsuite.Timestamp,
		TestsTotal: totalRunSuite + totalSkippedSuite,
		TestsRun: totalRunSuite,
		// TestsFailedSuite: totalFailedSuite,
		TestsSkipped: totalSkippedSuite,
		TestCases:     testCases,
	}
}

func WriteXMLToFile(report Report, filePath string) error {
	data, err := xml.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("Error while marshalling Report XML:\n %w", err)
	}
	return os.WriteFile(filePath, data, 0644)
}
