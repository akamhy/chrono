package main

import (
	"fmt"
	"os"
)

func PrintUsage() {
	// print the usage message
	fmt.Println("Usage: chrono [command] [options]")
	fmt.Println("Commands:")
	fmt.Println("  stopwatch, sw   Start the stopwatch")
	fmt.Println("  countdown, cd   Start the countdown timer")
	fmt.Println("  help            Show this help message")
	fmt.Println("Options:")
	fmt.Println("  -h, --help      Show this help message")

}

func StopwatchHelp() {
	fmt.Println("Usage: chrono stopwatch [options]")
	fmt.Println("Options:")
	fmt.Println("  -r, --resume   Resume the stopwatch from the last saved time")
	fmt.Println("  -h, --help     Show this help message")

	// exit the program
	os.Exit(0)

}

func CountdownTimerHelp() {
	fmt.Println("Usage: chrono countdown [options]")
	fmt.Println("Options:")
	fmt.Println("  [time]         The integer value of the time in seconds to count down from")
	fmt.Println("  -h, --help     Show this help message")

	// exit the program
	os.Exit(0)

}

func CheckHelpFlag() {
	// check if the help flag is passed as the first argument
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help") {
		PrintUsage()
		return

	}

	if len(os.Args) > 2 && (os.Args[2] == "-h" || os.Args[2] == "--help" || os.Args[2] == "help") {

		if os.Args[1] == "stopwatch" || os.Args[1] == "sw" {
			println("here")
			StopwatchHelp()
			return

		}

		if os.Args[1] == "countdown" || os.Args[1] == "cd" || os.Args[1] == "timer" {
			CountdownTimerHelp()
			return

		}
	}

}
