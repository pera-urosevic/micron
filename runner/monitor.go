package runner

import (
	"time"

	"micron/check"
	"micron/execute"
	"micron/system"
	"micron/types"
)

func Monitor(config *types.Config) {
	for i, job := range config.Monitor {
		if !job.Enabled {
			continue
		}
		if check.Running(job.Cmd) {
			continue
		}
		now := time.Now().Format(time.RFC3339)
		system.Log("Monitor |", job.Name, "[", job.LastRun, "=>", now, "]")
		execute.WithDetach(job.Name, job.Cmd, job.Args...)
		config.Monitor[i].LastRun = now
		config.Changed = true
	}
}
