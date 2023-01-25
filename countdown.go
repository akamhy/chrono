package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// function to act as a countdown timer
func CountdownTimer() {
	var timer int
	// check if the the timer argument is passed
	if len(os.Args) > 2 {
		// if the timer argument is passed, set the timer to the value of the argument
		fmt.Sscanf(os.Args[2], "%d", &timer)
	} else {
		// if the timer argument is not passed, ask the user to enter the timer value
		fmt.Print("Enter the timer value in seconds: ")
		fmt.Scanf("%d", &timer)

		// remove the previous line from the terminal
		fmt.Printf("\033[1A\033[2K")
	}

	// if the timer value is less than 1, exit the program
	if timer < 1 {
		// print that time is up
		fmt.Println("Time is up!")
		return
	}

	// hide the cursor when the timer is running
	_, _ = exec.Command("tput", "civis").Output()

	// show the cursor again when the program exits
	defer func() {
		_, _ = exec.Command("tput", "cnorm").Output()
	}()

	// print the remaining time and keep updating it
	// the timer will stop when the time is up or when the user presses Ctrl+C
	// clear the the last printed time when the timer is updated
	for {
		fmt.Print("\033[K")
		fmt.Printf("\r%d", timer)
		fmt.Print("\033[K")
		time.Sleep(time.Second)
		timer--
		if timer < 0 {
			fmt.Println()
			break
		}
	}

	// clear the the last printed line
	fmt.Print("\033[1A\033[2K")

	// print that time is up
	fmt.Println("Time is up!")

}
