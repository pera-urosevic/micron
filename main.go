package main

import (
	"log"

	"micron/admin"
	"micron/demo"
	"micron/storage"
	"micron/system"
	"micron/ticker"
)

func main() {
	system.Log("MiCron")

	err := storage.ConfigLoad()
	if err != nil {
		log.Fatal(err)
		return
	}

	if demo.Run() {
		return
	}

	system.Log("Commands: [P]ause, [E]dit, [Q]uit")
	storage.PrintConfig()

	go ticker.Start()
	admin.Start()
	system.Log("MiCron ended")
}
