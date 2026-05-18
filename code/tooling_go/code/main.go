package main

import "fmt"
import "flag"
import "strings"
import cfg "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/config"
import junit "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/jUnit"
import disc "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/testDiscovery"
import out "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/xmlOutput"

func main() {
	modifiedDiscoveryPath := flag.String("disc", "", "override test discovery path")
	modifiedJUnitXMLDirectory := flag.String("junit", "", "override junit xml directory path")
	modifiedOutputPath := flag.String("out", "", "override output path")
	flag.Parse()

	// Read config
	config, err := cfg.ReadConfig("./config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("config:\n", config, "\n")  // Debug

	discoveryCmd := config.TestDiscoveryPath
	if *modifiedDiscoveryPath != "" {
	parts := strings.Split(*modifiedDiscoveryPath, " ")
		discoveryCmd = cfg.Command{
			Command: parts[0],
			Args: parts[1:],
		}
	}

	jUnitXMLDirectory := config.JUnitXMLDirectory
	if *modifiedJUnitXMLDirectory != "" {
		jUnitXMLDirectory = *modifiedJUnitXMLDirectory
	}

	outputPath := config.OutputPath
	if *modifiedOutputPath != "" {
		outputPath = *modifiedOutputPath
	}

	// Read all existing tests from the user-configured script
	allSuites, err := disc.RunTestDiscoveryScript(discoveryCmd)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Suits from discovery:\n", allSuites, "\n")  // Debug

	// Read all tests in the JUnit XML output of the last run (if existing)
	allSuitesJUnit, err := junit.ReadJUnitTestSuites(jUnitXMLDirectory)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Suites from JUnit XML:\n", allSuitesJUnit, "\n")  // Debug

	report := out.MatchTests(allSuites, allSuitesJUnit)
	out.WriteXMLToFile(report, outputPath)
	fmt.Println("Successfully created report: \n", report)
}
