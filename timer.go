package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	// use this cross-platform notifier library to help
	// with making a simple solution
	"github.com/gen2brain/beeep"
)

func main() {
	// get all, but the first (program location), from cli arguments
	args := os.Args[1:]

	// convert the first argument to a 64-bit integer
	duration, parseErr := strconv.ParseInt(args[0], 10, 64)

	// if there's an issue parsing the string to an integer
	// handle the error here
	if parseErr != nil {
		panic(parseErr)
	}

	// join the remaining arguments with space delimiting
	// and assign to this title variable
	title := strings.Join(args[1:], " ")

	createTimer(duration, title)
}

func createTimer(duration int64, title string) {
	// give feedback to the user that something is working
	fmt.Println("timer set for", duration, "s for event \""+title+"\"")

	// convert the int64 to a item.Duration and cast it to
	// a time.Second just before sleeping for the given duration
	time.Sleep(time.Duration(duration) * time.Second)

	// notify the user with the title provided
	// todo: bundle the application and use the
	// last parameter to display an image
	err := beeep.Notify(title, "", "")
	if err != nil {
		panic(err)
	}
}
