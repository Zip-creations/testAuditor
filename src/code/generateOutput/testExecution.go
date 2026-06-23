package xmlOutput

import "fmt"
import "os"
import "os/exec"
import cfg "github.com/Zip-creations/optimize_CI_deterministic_builds/src/code/config"


func RunTestExecution(command cfg.Command , qualifiedNames []string) ([]byte, error) {
	args := make([]string, 0, len(command.Args)+len(qualifiedNames))
	args = append(args, command.Args...)
	args = append(args, qualifiedNames...)

	cmd := exec.Command(command.Command, args...)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("Error while running test execution script:\n %w", err)
	}

	return output, nil
}
