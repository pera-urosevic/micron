package demo

import (
	"micron/execute"
	"micron/storage"
	"micron/system"
	"os"
)

func Run() bool {
	if len(os.Args) < 3 {
		return false
	}
	name := os.Args[2]
	switch os.Args[1] {
	case "startup":
		for _, job := range storage.Config.Startup {
			if job.Name == name {
				system.Log("Demo:", job.Name, job.Cmd, job.Args)
				execute.WithOutput(job.Cmd, job.Args)
				return true
			}
		}
	case "monitor":
		for _, job := range storage.Config.Monitor {
			if job.Name == name {
				system.Log("Demo:", job.Name, job.Cmd, job.Args)
				execute.WithOutput(job.Cmd, job.Args)
				return true
			}
		}
	case "daily":
		for _, job := range storage.Config.Daily {
			if job.Name == name {
				system.Log("Demo:", job.Name, job.Cmd, job.Args)
				execute.WithOutput(job.Cmd, job.Args)
				return true
			}
		}
	case "weekly":
		for _, job := range storage.Config.Weekly {
			if job.Name == name {
				system.Log("Demo:", job.Name, job.Cmd, job.Args)
				execute.WithOutput(job.Cmd, job.Args)
				return true
			}
		}
	default:
		return false
	}
	return false
}
