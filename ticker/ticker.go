package ticker

import (
	"time"

	"micron/runner"
	"micron/storage"
)

var Running bool = true

func Start() {
	runner.Startup(&storage.Config)
	storage.ConfigSave()

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if Running {
					runner.Monitor(&storage.Config)
					runner.Daily(&storage.Config)
					runner.Weekly(&storage.Config)
					storage.ConfigSave()
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	for range ticker.C {
	}
}
