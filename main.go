package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type pomodoro interface {
	Doing(time.Duration) bool
}

// Gomodoro basic pomodoro structure
type Gomodoro struct {
	duration time.Duration
	rest     time.Duration
	resting  bool
}

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func callClear() {
	if v, ok := clear[runtime.GOOS]; ok {
		v()
	}
}

// Doing will return whether or not the pomodoro is running
func (g Gomodoro) Doing(duration time.Duration) bool {
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
			return true
		case <-ticker.C:
			callClear()
			if g.resting {
				fmt.Println("Relax, take it easy!")
			} else {
				fmt.Println("Doing a pomodoro, tackle those tasks!")
			}
			pomodoroTime++
			fmt.Println(secondsToMinutes(pomodoroTime))
			showProgressBar(duration, time.Duration(pomodoroTime)*time.Second)
		}
	}
}

func showProgressBar(pomodoroDuration time.Duration, actual time.Duration) {
	porcentLeft := int((actual * 100) / pomodoroDuration)
	fmt.Print("[")
	fmt.Print(strings.Repeat("=", porcentLeft))
	fmt.Print(">]")
}

func secondsToMinutes(inSeconds int) (str string) {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str = fmt.Sprintf("%02d:%02d", minutes, seconds)
	return
}

func main() {
	InitFlags()
}
