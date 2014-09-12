package main

import (
	"log"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	start := 232
	i := 0
	max := 16
	var delay time.Duration = 125

	for {
		for ; i <= max; i++ {
			colour := "colour" + strconv.Itoa(start+i)
			cmd := exec.Command("tmux", "set-window-option", "-g", "status-b", colour)
			err := cmd.Run()
			if err != nil {
				log.Println(err)
			}
			time.Sleep(delay * time.Millisecond)
		}

		time.Sleep(4 * delay * time.Millisecond)

		for ; i > 0; i-- {
			colour := "colour" + strconv.Itoa(start+i)
			cmd := exec.Command("tmux", "set-window-option", "-g", "status-b", colour)
			err := cmd.Run()
			if err != nil {
				log.Println(err)
			}
			time.Sleep(delay * time.Millisecond)
		}

		time.Sleep(4 * delay * time.Millisecond)
	}
}
