package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// function to act as a stopwatch
func Stopwatch() {

	var startTime int64
	var stopwatch string

	_, _ = exec.Command("tput", "civis").Output() // Hide the cursor

	defer func() {
		exec.Command("tput", "cnorm").Output() // Show the cursor again when the program exits
		os.Exit(0)
	}()

	// If the resume flag is passed, resume the stopwatch from the last saved time
	if len(os.Args) > 2 && (os.Args[2] == "-r" ||
		os.Args[2] == "--resume" ||
		os.Args[2] == "-resume" ||
		os.Args[2] == "--r" ||
		os.Args[2] == "resume" ||
		os.Args[2] == "r") {
		if _, err := os.Stat(os.Getenv("HOME") + "/.chrono"); os.IsNotExist(err) {
			startTime = time.Now().Unix()
			f, _ := os.Create(os.Getenv("HOME") + "/.chrono")
			defer f.Close()
			f.WriteString(fmt.Sprintf("%d", startTime))
		} else {
			f, _ := os.Open(os.Getenv("HOME") + "/.chrono")
			defer f.Close()
			fmt.Fscanf(f, "%d", &startTime)
		}
	} else {
		startTime = time.Now().Unix()
		f, _ := os.Create(os.Getenv("HOME") + "/.chrono")
		defer f.Close()
		f.WriteString(fmt.Sprintf("%d", startTime))
	}

	for {
		stopwatch = time.Now().UTC().Add(-time.Duration(startTime) * time.Second).Format("15:04:05")
		fmt.Printf("\r%s", stopwatch)
		time.Sleep(30 * time.Millisecond)
	}
}
