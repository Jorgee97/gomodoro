package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	duration     = flag.Duration("duration", time.Second*10, "Define the time of the pomodoro.")
	restDuration = flag.Duration("rest", time.Minute*5, "Define the rest duration after every pomodoro.")
)

// InitFlags initialize the flags for the pomodoro.
func InitFlags() {
	flag.Parse()

	gomodoro := &Gomodoro{
		duration:      *duration,
		rest:          *restDuration,
		notify:        false,
		soundLocation: "",
	}

	doing := gomodoro.Doing(gomodoro.duration, gomodoro.notify, gomodoro.soundLocation)
	if !doing {
		fmt.Println("Error while running the pomodoro.")
	}
}
