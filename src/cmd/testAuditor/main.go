package main

import "fmt"
import "os"
import "io"
import rep "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/report"
import disc "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/testDiscovery"
import out "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/matchTests"
import input "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/parseInput"

func main() {
	stdInput, err := io.ReadAll(os.Stdin)
	if err != nil {
		// write on Stderr, since Stdout is the expected channel for the produced XML
		fmt.Fprintln(os.Stderr, "Failed to read stdin:", err)
		os.Exit(1)
	}
	// Parse stdin
	parsedInput, err := input.ParseInput(stdInput)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse input:", err)
		os.Exit(1)
	}
	// fmt.Println("ParseInput result:", parsedInput)  // Debug

	// Parse all existing tests from input
	allSuites, err := disc.ParseTestDiscovery(parsedInput.TestDiscovery.Content)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// 0 tests in test discovey means there are no tests expected in the project
	if len(allSuites.DiscoveryTestcases) == 0 {
		fmt.Fprintln(os.Stderr, "Test discovery resulted in 0 tests found. Aborting.")
		os.Exit(0)
	}
	// fmt.Println("ParseTestDiscovery result:", allSuites)  // Debug

	// Parse all existing reports from input
	existingReports, err := rep.ParseJUnitTestSuites(parsedInput.Reports)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// fmt.Println("ParseJUnitTestSuites result:", existingReports)  // Debug

	result := out.MatchTests(allSuites, existingReports)
	// Write result to stdout
	fmt.Println(result)
}
