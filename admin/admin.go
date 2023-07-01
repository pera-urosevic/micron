package admin

import (
	"bufio"
	"os"

	"dreamsleep.cloud/micron/storage"
	"dreamsleep.cloud/micron/ticker"
)

func Start() {
	for {
		reader := bufio.NewReader(os.Stdin)
		char, _ := reader.ReadByte()
		switch char {
		case 'q':
			os.Exit(0)
		case 'p':
			ticker.Running = !ticker.Running
		case 'e':
			ticker.Running = false
			storage.ConfigEdit()
			ticker.Running = true
		}
	}
}
