package check

import (
	"testing"
)

func TestCheckRunning(t *testing.T) {
	running := Running("C:\\NotReal\\explorer.exe")
	if !running {
		t.Errorf("Check running lowercase failed")
	}

	running = Running("C:\\NotReal\\EXPLORER.exe")
	if !running {
		t.Errorf("Check running uppercase failed")
	}

	running = Running("C:\\NotReal\\AlsoNotReal.exe")
	if running {
		t.Errorf("Check running non existing failed")
	}
}
