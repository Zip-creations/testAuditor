package main

import "fmt"
import "flag"
import "strings"
import cfg "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/config"
import junit "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/jUnit"
import disc "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/testDiscovery"
import out "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/xmlOutput"

func main() {
	dummyGitNote := `<?xml version="1.0" encoding="utf-8"?>
	<testsuites name="pytest tests">
		<testsuite name="pytest" errors="0" failures="0" skipped="1" tests="4" time="0.013"
			timestamp="2026-05-12T13:36:20.355444+02:00" hostname="DESKTOP-I9APAL0">
			<testcase classname="test.test_simple" name="test_add_item" time="0.000">
				<properties>
					<property name="example_key" value="1" />
				</properties>
			</testcase>
			<testcase classname="test.test_simple" name="test_removing_items" time="0.000" />
			<testcase classname="test.test_simple" name="test_skipping" time="0.000">
				<skipped type="pytest.skip"
					message="this test will be skipped, to see how pytest handles skipped tests in JUnit XML reports">/home/ruben/dev/bachelor/project/code/sample_projects/python/code/test/test_simple.py:35:
					this test will be skipped, to see how pytest handles skipped tests in JUnit XML
					reports</skipped>
			</testcase>
			<testcase classname="test.test_identity" name="test_copy" time="0.000" />
		</testsuite>
	</testsuites>
	
	<?xml version="1.0" encoding="utf-8"?>
	<testsuites name="pytest tests">
		<testsuite name="pytest" errors="0" failures="0" skipped="1" tests="4" time="0.013"
			timestamp="2026-05-12T13:36:20.355444+02:00" hostname="DESKTOP-I9APAL0">
			<testcase classname="test.test_simple" name="test_add_item" time="0.000">
				<properties>
					<property name="example_key" value="1" />
				</properties>
			</testcase>
			<testcase classname="test.test_simple" name="test_removing_items" time="0.000" />
			<testcase classname="test.test_simple" name="test_skipping" time="0.000">
				<skipped type="pytest.skip"
					message="this test will be skipped, to see how pytest handles skipped tests in JUnit XML reports">/home/ruben/dev/bachelor/project/code/sample_projects/python/code/test/test_simple.py:35:
					this test will be skipped, to see how pytest handles skipped tests in JUnit XML
					reports</skipped>
			</testcase>
			<testcase classname="test.test_identity" name="test_copy" time="0.000" />
		</testsuite>
	</testsuites>
	
	<?xml version='1.0' encoding='UTF-8'?>
	<testsuite hostname="DESKTOP-I9APAL0" name="sample_projects.scala.code.test.TestAddingItems" tests="1" errors="0" failures="0" skipped="0" time="0.068" timestamp="2026-06-15T16:29:44">
			  <properties>
		  <property name="java.specification.version" value="21"/><property name="sun.jnu.encoding" value="UTF-8"/><property name="sun.arch.data.model" value="64"/><property name="java.vendor.url" value="https://openjdk.org/"/><property name="sun.boot.library.path" value="/nix/store/kiaiabxhnl2a15zpiamqvg13wfx1ikal-openjdk-21.0.10+7/lib/openjdk/lib"/><property name="sun.java.command" value="/nix/store/bpvi6p1irywz92kw17gyp3xq64gcmcyb-sbt-1.12.4/share/sbt/bin/sbt-launch.jar clean test"/><property name="jdk.debug" value="release"/><property name="java.specification.vendor" value="Oracle Corporation"/><property name="java.version.date" value="2026-01-20"/><property name="java.home" value="/nix/store/kiaiabxhnl2a15zpiamqvg13wfx1ikal-openjdk-21.0.10+7/lib/openjdk"/><property name="file.separator" value="/"/><property name="java.vm.compressedOopsMode" value="32-bit"/><property name="line.separator" value="
	"/><property name="java.vm.specification.vendor" value="Oracle Corporation"/><property name="java.specification.name" value="Java Platform API Specification"/><property name="sun.management.compiler" value="HotSpot 64-Bit Tiered Compilers"/><property name="java.runtime.version" value="21.0.10+7-nixos"/><property name="user.name" value="ruben"/><property name="log4j.ignoreTCL" value="true"/><property name="file.encoding" value="UTF-8"/><property name="jnidispatch.path" value="/home/ruben/.cache/JNA/temp/jna15134823488361060376.tmp"/><property name="jna.loaded" value="true"/><property name="java.io.tmpdir" value="/tmp"/><property name="java.version" value="21.0.10"/><property name="java.vm.specification.name" value="Java Virtual Machine Specification"/><property name="native.encoding" value="UTF-8"/><property name="java.library.path" value="/usr/java/packages/lib:/usr/lib64:/lib64:/lib:/usr/lib"/><property name="stderr.encoding" value="UTF-8"/><property name="java.vendor" value="N/A"/><property name="scala.ext.dirs" value="/home/ruben/.sbt/1.0/java9-rt-ext-n_a_21_0_10"/><property name="sun.io.unicode.encoding" value="UnicodeLittle"/><property name="java.class.path" value="/nix/store/bpvi6p1irywz92kw17gyp3xq64gcmcyb-sbt-1.12.4/share/sbt/bin/sbt-launch.jar"/><property name="java.vm.vendor" value="Oracle Corporation"/><property name="jline.shutdownhook" value="false"/><property name="user.timezone" value="Europe/Berlin"/><property name="java.vm.specification.version" value="21"/><property name="os.name" value="Linux"/><property name="sun.java.launcher" value="SUN_STANDARD"/><property name="sun.cpu.endian" value="little"/><property name="user.home" value="/home/ruben"/><property name="user.language" value="en"/><property name="sbt.script" value="/nix/store/bpvi6p1irywz92kw17gyp3xq64gcmcyb-sbt-1.12.4/bin/sbt"/><property name="sbt.ipcsocket.tmpdir" value="/run/user/1000/.sbt9ad88350/ipcsocket"/><property name="swoval.tmpdir" value="/run/user/1000/.sbt9ad88350/swoval"/><property name="jline.esc.timeout" value="0"/><property name="stdout.encoding" value="UTF-8"/><property name="path.separator" value=":"/><property name="os.version" value="6.6.87.2-microsoft-standard-WSL2"/><property name="jna.nosys" value="true"/><property name="java.runtime.name" value="OpenJDK Runtime Environment"/><property name="java.vm.name" value="OpenJDK 64-Bit Server VM"/><property name="jna.platform.library.path" value="/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu:/usr/lib64:/lib64:/usr/lib:/lib:/usr/lib/x86_64-linux-gnu/libfakeroot:/usr/lib/wsl/lib"/><property name="java.vendor.url.bug" value="https://bugreport.java.com/bugreport/"/><property name="user.dir" value="/home/ruben/dev/bachelor/project/code/code/sample_projects/scala/code"/><property name="os.arch" value="amd64"/><property name="java.vm.info" value="mixed mode, sharing"/><property name="java.vm.version" value="21.0.10+7-nixos"/><property name="java.class.version" value="65.0"/>
		</properties>
			  <testcase classname="sample_projects.scala.code.test.TestAddingItems" name="testTodo()" time="0.068">
						  
						</testcase>
			  <system-out><![CDATA[]]></system-out>
			  <system-err><![CDATA[]]></system-err>
			</testsuite>`  // test
	
	modifiedDiscoveryPath := flag.String("disc", "", "override test discovery path")
	modifiedExecutionPath := flag.String("out", "", "override output path")
	contentString := flag.String("c", dummyGitNote, "provide JUnit XML content")  // TODO: remove default
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

	executionCmd := config.TestExecutionPath
	if *modifiedExecutionPath != "" {
	parts := strings.Split(*modifiedExecutionPath, " ")
		executionCmd = cfg.Command{
			Command: parts[0],
			Args: parts[1:],
		}
	}

	// Read all existing tests from the user-configured script
	allSuites, err := disc.RunTestDiscoveryScript(discoveryCmd)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Suits from discovery:\n", allSuites, "\n")  // Debug

	// Read all tests in the JUnit XML output of the last run (if existing)
	allSuitesJUnit, err := junit.ReadGitNote(*contentString)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Suites from JUnit XML:\n", allSuitesJUnit, "\n")  // Debug

	report := out.MatchTests(allSuites, allSuitesJUnit)
	if len(report) == 0 {
		return
	}
	executionErr := out.RunTestScript(executionCmd, report)
	if executionErr != nil {
		fmt.Println(executionErr)
		return
	}
	fmt.Println("Successfully created report: \n", report)
}
