package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"time"
)

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is an implementation of Sleeper with a defined delay.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep will pause execution for the defined Duration.
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper.
func Countdown(out io.Writer, sleeper Sleeper) error {
	for i := range countDownFrom(3) {
		if _, err := fmt.Fprintln(out, i); err != nil {
			return err
		}
		sleeper.Sleep()
	}

	if _, err := fmt.Fprint(out, finalWord); err != nil {
		return err
	}
	return nil
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	if err := Countdown(os.Stdout, sleeper); err != nil {
		panic(err)
	}
}
