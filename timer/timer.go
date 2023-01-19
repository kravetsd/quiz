package timer

import (
	"fmt"
	"time"
)

func StartTimer() {
	fmt.Println("Starting timer...")
	fmt.Println(time.Now().Format("15:04:05"))
}
