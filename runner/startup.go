package runner

import (
	"time"

	"dreamsleep.cloud/micron/check"
	"dreamsleep.cloud/micron/execute"
	"dreamsleep.cloud/micron/system"
	"dreamsleep.cloud/micron/types"
)

func Startup(config *types.Config) {
	for i, job := range config.Startup {
		if !job.Enabled {
			continue
		}
		if check.Running(job.Cmd) {
			continue
		}
		now := time.Now().Format(time.RFC3339)
		system.Log("Startup |", job.Name, "[", job.LastRun, "=>", now, "]")
		execute.WithDetach(job.Name, job.Cmd, job.Args...)
		config.Startup[i].LastRun = now
		config.Changed = true
	}
}
