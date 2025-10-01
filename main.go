package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) > 3 {
		oph := os.Args[1]
		opm := os.Args[2]
		ops := os.Args[3]
		hours, err := strconv.Atoi(oph)
		if err != nil {
			fmt.Println("Invalid hours : enter number")
			return
		}
		minutes, err := strconv.Atoi(opm)
		if err != nil {
			fmt.Println("Invalid minutes : enter number")
			return
		}
		seconds, err := strconv.Atoi(ops)
		if err != nil {
			fmt.Println("Invalid seconds : enter number")
			return
		}
		run(hours, minutes, seconds)
	} else {
		fmt.Println("Usage: <hours> <minutes> <seconds>")
	}
}

func run(hours, minutes, seconds int) {
	for {
		time.Sleep(time.Second * 1)
		seconds--
		if seconds < 0 {
			seconds = 59
			minutes--
		}
		if minutes < 0 {
			minutes = 59
			hours--
		}
		if hours < 0 {
			hours = 23
		}
		if seconds == 0 && minutes == 0 && hours == 0 {
			fmt.Println("timeup")
			break
		}
		fmt.Printf("%02d:%02d:%02d\r", hours, minutes, seconds)
	}
}
