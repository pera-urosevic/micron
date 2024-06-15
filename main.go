package main

import (
	"log"

	"micron/admin"
	"micron/storage"
	"micron/system"
	"micron/ticker"
)

func main() {
	system.Log("MiCron")

	system.Log("Commands: [P]ause, [E]dit, [Q]uit")

	err := storage.ConfigLoad()
	if err != nil {
		log.Fatal(err)
		return
	}

	go ticker.Start()
	admin.Start()
	system.Log("MiCron ended")
}
