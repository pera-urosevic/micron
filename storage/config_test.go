package storage

import (
	"testing"
)

func TestConfigPath(t *testing.T) {
	path, err := configPath()
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if len(path) == 0 {
		t.Errorf("path is empty")
	}
}

func TestConfigLoad(t *testing.T) {
	err := ConfigLoad()
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if len(Config.Startup) == 0 {
		t.Errorf("startup config is empty")
	}
	if len(Config.Monitor) == 0 {
		t.Errorf("monitor config is empty")
	}
	if len(Config.Daily) == 0 {
		t.Errorf("daily config is empty")
	}
	if len(Config.Weekly) == 0 {
		t.Errorf("weekly config is empty")
	}
}

func TestConfigSave(t *testing.T) {
	err := ConfigLoad()
	if err != nil {
		t.Errorf("err: %v", err)
	}
	err = ConfigSave()
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
