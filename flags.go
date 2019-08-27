package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	repeat       = flag.Int("repeat", 1, "Define how many pomodoros you want to make")
	duration     = flag.Duration("duration", time.Second*10, "Define the time of the pomodoro.")
	restDuration = flag.Duration("rest", time.Second*5, "Define the rest duration after every pomodoro.")
)

// InitFlags initialize the flags for the pomodoro.
func InitFlags() {
	flag.Parse()

	gomodoro := &Gomodoro{
		duration: *duration,
		rest:     *restDuration,
		resting:  false,
	}

	rest := *repeat
	for i := 0; i < rest; i++ {
		doing := gomodoro.Doing(gomodoro.duration)
		if !doing {
			fmt.Println("Error while running the pomodoro.")
		}
		if *repeat > 1 && doing {
			gomodoro.resting = true
			doing = gomodoro.Doing(gomodoro.rest)
			if doing {
				gomodoro.resting = false
			}

			// I decrement repeat because the rest after the last pomodoro is kind of unnecesary.
			*repeat--
		}
	}
}
