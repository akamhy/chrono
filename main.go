package main

import (
	"os"
)

func main() {

	// If the help flag is passed, print the usage message and exit
	CheckHelpFlag()

	// if the first argument is "stopwatch" or "sw", run the stopwatch
	if len(os.Args) > 1 && (os.Args[1] == "stopwatch" || os.Args[1] == "sw") {
		Stopwatch()
		// exit the program
		os.Exit(0)
	}

	// if the first argument is "countdown" or "cd", run the countdown timer
	if len(os.Args) > 1 && (os.Args[1] == "countdown" || os.Args[1] == "cd" || os.Args[1] == "timer") {
		CountdownTimer()
		// exit the program
		os.Exit(0)
	}

	// if no arguments are passed, show the local date time and keep updating it
	if len(os.Args) == 1 {
		ShowTime()
		// exit the program
		os.Exit(0)
	}

	// exit the program
	os.Exit(0)

}
