package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type RealSleeper struct {
	duration time.Duration
}

func (s *RealSleeper) Sleep() {
	time.Sleep(s.duration)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	const finalWord = "GO!"
	const startNum = 3
	for i := startNum; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	realSleeper := RealSleeper{duration: 1 * time.Second}
	Countdown(os.Stdout, &realSleeper)
}
