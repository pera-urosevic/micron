package execute

import (
	"testing"
)

func TestCommandExe(t *testing.T) {
	exe := GetExe("C:\\NotReal\\go.exe")
	if exe != "go.exe" {
		t.Errorf("Command exe failed")
	}
}
