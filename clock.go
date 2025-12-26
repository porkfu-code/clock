package clock

import (
	"fmt"
	"log"
	"strings"
	"time"
)

/*
Clock simply displays the current time in the format HH:MM:SS
*/

type clock struct {
	tickRate  int // the clock will tick at this rate in Hz
	showColon bool
}

func NewClock() *clock {
	return &clock{
		tickRate:  2,
		showColon: true,
	}
}

func (c *clock) displayTime() error {
	now := time.Now()

	formattedNow := now.Format("15:04:05")

	timeSegments := strings.Split(formattedNow, ":")
	if len(timeSegments) < 3 {
		return fmt.Errorf("Got time %s and failed to split hours, minutes, seconds on ':'", formattedNow)
	}

	hours := timeSegments[0]
	minutes := timeSegments[1]
	seconds := timeSegments[2]

	if c.showColon {
		fmt.Printf("\r%s:%s:%s", hours, minutes, seconds)
	} else {
		fmt.Printf("\r%s %s %s", hours, minutes, seconds)
	}

	c.showColon = !c.showColon // oscillate between showing and not showing ':'

	return nil
}

func (c *clock) wait() {
	sleepDuration := time.Duration(float64(time.Second) * (1.0 / float64(c.tickRate)))

	time.Sleep(sleepDuration)
}

func (c *clock) Run() {
	for {
		if err := c.displayTime(); err != nil {
			log.Fatalln("clock.displayTime:", err)
			return
		}

		c.wait()
	}
}
