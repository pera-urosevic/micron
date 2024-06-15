package system

import (
	"fmt"
	"time"
)

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func Log(args ...any) {
	timestamp := fmt.Sprintf("[%s]", TimeFormat(time.Now()))
	line := append([]any{timestamp}, args...)
	fmt.Println(line...)
}
