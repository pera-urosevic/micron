package system

import (
	"fmt"
	"time"
)

func Log(args ...any) {
	timestamp := fmt.Sprintf("[%s]", time.Now().Format("2006-01-02 15:04:05"))
	line := append([]any{timestamp}, args...)
	fmt.Println(line...)
}
