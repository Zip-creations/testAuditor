package xmlOutput

import "fmt"
import "os"
import "os/exec"
import cfg "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/config"


func RunTestScript(command cfg.Command , qualifiedNames []string) error {
	out, err := exec.Command(command.Command, command.Args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error executing test discovery script: %w\n%s", err, out)
	}
	args := make([]string, 0, len(command.Args)+len(qualifiedNames))
	args = append(args, command.Args...)
	args = append(args, qualifiedNames...)

	cmd := exec.Command(command.Command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("test execution script failed: %w", err)
	}

	return nil
}
