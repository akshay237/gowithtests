package gomocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countDownStart = 3
	finalWord      = "Go!"
	write          = "write"
	sleep          = "sleep"
)

// func CountDown(writer *bytes.Buffer) {
// 	fmt.Fprint(writer, "3")
// }

// Make the countdown generic use io.writer instead of bytes.buffer
// If we want to pass the func of a interface them inject the interfcae ans we can pass all the struct that implments methods of that interface

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i = i - 1 {
		sleeper.Sleep()
	}

	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, finalWord)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleep struct{}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (d *DefaultSleep) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpyOperationsTrack struct {
	calls []string
}

func (s *SpyOperationsTrack) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *SpyOperationsTrack) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return
}

// Configure the sleeper
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// Spies are a kind of mock which can record how a dependency is used.
//They can record the arguments sent in, how many times it has been called, etc.

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}
