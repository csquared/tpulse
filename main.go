package main

import (
	"log"
	"os/exec"
	"strconv"
	"time"
)

var delay time.Duration = 125

func runAndSleep(colour string) {
	cmd := exec.Command("tmux", "set-window-option", "-gq", "status-b", colour)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	time.Sleep(delay * time.Millisecond)
}

func main() {
	start := 232
	i := 0
	max := 16

	for {
		for ; i <= max; i++ {
			runAndSleep("colour" + strconv.Itoa(start+i))
		}

		time.Sleep(4 * delay * time.Millisecond)

		for ; i > 0; i-- {
			runAndSleep("colour" + strconv.Itoa(start+i))
		}

		time.Sleep(4 * delay * time.Millisecond)
	}
}
