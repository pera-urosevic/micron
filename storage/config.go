package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"micron/execute"
	"micron/system"
	"micron/types"
)

var Config types.Config

func configPath() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\\micron.json", dirname), nil
}

func PrintConfig() {
	for _, item := range Config.Startup {
		system.Log("Startup:", item.Name)
	}

	for _, item := range Config.Monitor {
		system.Log("Monitor:", item.Name)
	}

	for _, item := range Config.Daily {
		system.Log("Daily:", item.Name)
	}

	for _, item := range Config.Weekly {
		system.Log("Weekly:", item.Name)
	}
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

func ConfigEdit() error {
	path, err := configPath()
	if err != nil {
		return err
	}

	_, err = execute.WithOutput(Config.Editor, []string{path})
	if err != nil {
		return err
	}

	return ConfigLoad()
}
