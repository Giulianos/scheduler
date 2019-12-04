package main

import (
	"fmt"
	"time"

	"github.com/Giulianos/scheduler"
)

func main() {
	s := scheduler.New()

	// Schedule a job to run every minute
	s.Schedule(
		func() {
			fmt.Println("Every minute job")
		},
		scheduler.CronRule{
			Minute:  -1,
			Hour:    14,
			Day:     -1,
			Month:   -1,
			Weekday: -1,
		},
	)

	// Schedule a job to run once (1min after execution)
	nextMoment := time.Now().Add(1 * time.Minute)
	s.Schedule(
		func() {
			fmt.Println("Single execution job")
		},
		scheduler.CronRule{
			Minute:  nextMoment.Minute(),
			Hour:    nextMoment.Hour(),
			Day:     nextMoment.Day(),
			Month:   nextMoment.Month(),
			Weekday: nextMoment.Weekday(),
		},
	)

	// Running scheduler
	s.Run()
	for {

	}
}
