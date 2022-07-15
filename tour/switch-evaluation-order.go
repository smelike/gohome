package main

import (
	"fmt"
	"time"
)

// evaluate cases from top to bottom, stopping when a case succeeds.
// [Question]: what does it mean switch cases the values involved need not be integer?
// switch i?
func switch_order_main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(time.Friday + 1) // Saturday
	switch time.Saturday {
	case today + 0:
		fmt.Println("today.", today)
	case today + 1:
		fmt.Println("tomorrow.", today)
	case today + 2:
		fmt.Println("in two days.", today)
	default:
		fmt.Println("Too far away.")
	}
}
