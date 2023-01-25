package main

import (
	"fmt"
	"time"
)

// function to show the current local date time and keep updating it
func ShowTime() {
	// infinite loop
	for {

		// get the current local date time
		t := time.Now()

		// print the current local date time
		fmt.Println(t.Format("Mon Jan 2 15:04:05 MST 2006"))

		// sleep for 1 second
		time.Sleep(time.Second)

		// clear the the last printed line
		fmt.Print("\033[1A\033[2K")
	}
}
