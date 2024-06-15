package check

import (
	"fmt"
	"strings"

	"micron/execute"
)

func Running(cmd string) bool {
	// run tasklist
	exe := execute.GetExe(cmd)
	tasklistCmd := "tasklist"
	tasklistArgs := []string{"/fo", "csv", "/fi", "sessionname eq Console", "/fi", fmt.Sprintf("imagename eq %s", exe)}
	output, err := execute.WithOutput(tasklistCmd, tasklistArgs)
	if err != nil {
		return true
	}
	// check if found
	outputLower := strings.ToLower(output)
	imageName := fmt.Sprintf("\"%s\"", exe)
	imageNameLower := strings.ToLower(imageName)
	running := strings.Contains(outputLower, imageNameLower)
	return running
}
