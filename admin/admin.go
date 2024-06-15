package admin

import (
	"bufio"
	"os"

	"micron/storage"
	"micron/system"
	"micron/ticker"
)

func Start() {
	for {
		reader := bufio.NewReader(os.Stdin)
		char, _ := reader.ReadByte()
		switch char {
		case 'q':
			system.Log("Quiting...")
			os.Exit(0)
		case 'p':
			ticker.Running = !ticker.Running
			if ticker.Running {
				system.Log("Resumed")
			} else {
				system.Log("Paused")
			}
		case 'e':
			ticker.Running = false
			system.Log("Editing config")
			storage.ConfigEdit()
			system.Log("Resumed")
			ticker.Running = true
		}
	}
}
