package main

import "fmt"
import "os"
import "io"
import rep "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/report"
import disc "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/testDiscovery"
import out "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/matchTests"
import iparse "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/parseInput"

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read stdin:", err)
		os.Exit(1)
	}
	parsedInput, err := iparse.ReadInput(string(input))
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to parse input:", err)
		os.Exit(1)
	}
	fmt.Println("ReadInput result:", parsedInput)

	// Read all existing tests from the user-configured script
	allSuites, err := disc.XMLtoDiscoveryTestsuite(parsedInput.TestDiscovery.Content)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(allSuites.DiscoveryTestcases) == 0 {
		// write on Stderr, since Stdout is the expected route for the produced XML
		fmt.Fprintln(os.Stderr, "test discovery resulted in 0 tests found. Aborting.")
		os.Exit(0)
	}
	fmt.Println("ReadDiscovery result:", allSuites)

	existingReports, err := rep.ReadJUnitTestSuites(parsedInput.Reports)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("ReadJUnitTestSuites result:", existingReports)

	report := out.MatchTests(allSuites, existingReports)
	if len(report) == 0 {
		// write on Stderr, since Stdout is the expected route for the produced XML
		fmt.Fprintln(os.Stderr, "All discovered tests have already been executed.\nTerminating.")
		os.Exit(0)
	}
	fmt.Println("report: ", report)
}
