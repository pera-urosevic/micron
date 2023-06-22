package runner

import (
	"time"

	"dreamsleep.cloud/micron/check"
	"dreamsleep.cloud/micron/execute"
	"dreamsleep.cloud/micron/system"
	"dreamsleep.cloud/micron/types"
)

func Weekly(config *types.Config) {
	for i, job := range config.Weekly {
		if !job.Enabled {
			continue
		}
		if !check.WeekDayTime(job.Day, job.Time, job.LastRun) {
			continue
		}
		now := time.Now().Format(time.RFC3339)
		system.Log("Weekly |", job.Name, "[", job.LastRun, "=>", now, "]")
		execute.WithDetach(job.Name, job.Cmd, job.Args...)
		config.Weekly[i].LastRun = now
		config.Changed = true
	}
}
