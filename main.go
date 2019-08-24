package main

import (
	"fmt"
	"time"
)

type pomodoro interface {
	Doing(time.Duration, bool, string) bool
	Resting() bool
}

type Gomodoro struct {
	duration      time.Duration
	rest          time.Duration
	notify        bool
	soundLocation string
}

// Doing will return whether or not the pomodoro is running
func (g Gomodoro) Doing(duration time.Duration, notify bool, soundLocation string) bool {
	pomodoroTime := 0

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(duration)
		done <- true
	}()
	for {
		select {
		case <-done:
			// Should call rest.
			return true
		case t := <-ticker.C:
			pomodoroTime++
			fmt.Println(secondsToMinutes(pomodoroTime))
			fmt.Println("Current time: ", t)
		}
	}
}

func secondsToMinutes(inSeconds int) (str string) {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str = fmt.Sprintf("%d:%d", minutes, seconds)
	return
}

func main() {
	InitFlags()
}
