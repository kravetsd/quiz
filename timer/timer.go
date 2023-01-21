package timer

import (
	"fmt"
	"time"
)

func StartTimer() {
	fmt.Println("Starting timer...")
	start := time.Now()
	fmt.Println("Timer started at: ", start)
	time.Sleep(5 * time.Second)
	end := time.Now()
	fmt.Println("Timer ended at: ", end)
	fmt.Println("Time elapsed: ", end.Sub(start))
	// stating timer in parallel to quiz questions
	// good to use time.Duration
	// good to ask user to set duaraion of the quiz in the flag like 1ms, 1s,1m 1h and so on via time.ParseDuration

}
