package runner

import (
	"time"

	"micron/check"
	"micron/execute"
	"micron/system"
	"micron/types"
)

func Daily(config *types.Config) {
	for i, job := range config.Daily {
		if !job.Enabled {
			continue
		}
		if !check.DayTime(job.Time, job.LastRun) {
			continue
		}
		now := time.Now().Format(time.RFC3339)
		system.Log("Daily |", job.Name, "[", job.LastRun, "=>", now, "]")
		execute.WithDetach(job.Name, job.Cmd, job.Args...)
		config.Daily[i].LastRun = now
		config.Changed = true
	}
}
