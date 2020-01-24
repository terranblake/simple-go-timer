package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	args := os.Args[1:]
	amountOfTime, parseErr := strconv.ParseInt(args[0], 10, 64)
	if parseErr != nil {
		panic(parseErr)
	}

	title := strings.Join(args[1:], " ")

	fmt.Println("i will notify you of \"", title, "\" in", amountOfTime, "seconds")
	time.Sleep(time.Duration(amountOfTime) * time.Second)

	err := beeep.Notify(title, "", "")
	if err != nil {
		panic(err)
	}
}
