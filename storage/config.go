package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"dreamsleep.cloud/micron/system"
	"dreamsleep.cloud/micron/types"
)

var Config types.Config

func configPath() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\\micron.json", dirname), nil
}

func ConfigLoad() error {
	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Config)
	if err != nil {
		return err
	}

	system.Log("Startup:", len(Config.Startup))
	system.Log("Monitor:", len(Config.Monitor))
	system.Log("Daily:", len(Config.Daily))
	system.Log("Weekly:", len(Config.Weekly))

	Config.Changed = false
	return nil
}

func ConfigSave() error {
	if !Config.Changed {
		return nil
	}

	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(Config, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	Config.Changed = false
	return nil
}
