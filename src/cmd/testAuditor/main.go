package main

import "fmt"
import "flag"
import "strings"
import "os"
import "io"
import cfg "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/config"
import junit "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/jUnit"
import disc "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/testDiscovery"
import out "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/generateOutput"

func main() {
	modifiedDiscoveryPath := flag.String("disc", "", "override test discovery path")
	modifiedExecutionPath := flag.String("exec", "", "override test execution path")
	flag.Parse()

	// Read config
	config, err := cfg.ReadConfig("./config.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read stdin:", err)
		os.Exit(1)
	}

	discoveryCmd := config.TestDiscoveryPath
	if *modifiedDiscoveryPath != "" {
	parts := strings.Split(*modifiedDiscoveryPath, " ")
		discoveryCmd = cfg.Command{
			Command: parts[0],
			Args: parts[1:],
		}
	}

	executionCmd := config.TestExecutionPath
	if *modifiedExecutionPath != "" {
	parts := strings.Split(*modifiedExecutionPath, " ")
		executionCmd = cfg.Command{
			Command: parts[0],
			Args: parts[1:],
		}
	}

	// Read all existing tests from the user-configured script
	allSuites, err := disc.RunTestDiscovery(discoveryCmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if len(allSuites.DiscoveryTestcases) == 0 {
		// write on Stderr, since Stdout is the expected route for the produced XML
		fmt.Fprintln(os.Stderr, "test discovery resulted in 0 tests found. Aborting.")
		os.Exit(0)
	}

	// Read all tests in the JUnit XML output of the last run (if existing)
	allSuitesJUnit, err := junit.ReadGitNote(string(input))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	report := out.MatchTests(allSuites, allSuitesJUnit)
	if len(report) == 0 {
		// write on Stderr, since Stdout is the expected route for the produced XML
		fmt.Fprintln(os.Stderr, "All discovered tests have already been executed.\nTerminating.")
		os.Exit(0)
	}

	// Try to launch test execution script. If successfull, print the output so the hook can pick it up from stdout
	output, executionErr := out.RunTestExecution(executionCmd, report)
	if executionErr != nil {
		fmt.Fprintln(os.Stderr, executionErr)
		os.Exit(1)
	}
	fmt.Println(string(output))
}
