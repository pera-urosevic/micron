package execute

import (
	"os/exec"
	"strings"
	"time"

	"dreamsleep.cloud/micron/system"
)

func GetExe(cmd string) string {
	parts := strings.Split(cmd, "\\")
	exe := parts[len(parts)-1]
	return exe
}

func WithOutput(cmd string, args []string) (string, error) {
	output, err := exec.Command(cmd, args...).Output()
	return string(output), err
}

func WithDetach(name string, cmd string, args ...string) {
	time.AfterFunc(10*time.Second, func() {
		app := exec.Command(cmd, args...)
		err := app.Start()
		if err != nil {
			system.Log("Error executing with detach", err)
		} else {
			app.Process.Release()
			cmdline := cmd + " " + strings.Join(args, " ")
			system.Log("Execute |", cmdline)
		}
	})
}
