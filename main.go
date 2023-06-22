package main

import (
	"log"

	"dreamsleep.cloud/micron/storage"
	"dreamsleep.cloud/micron/system"
	"dreamsleep.cloud/micron/ticker"
)

func main() {
	system.Log("MiCron starting")

	err := storage.ConfigLoad()
	if err != nil {
		log.Fatal(err)
		return
	}

	ticker.Start()
	system.Log("MiCron ended")
}
