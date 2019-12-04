package main

import (
	"log"
	"time"

	"github.com/Giulianos/scheduler"
)

func main() {
	s := scheduler.New()

	// Schedule a job to run every minute
	j1 := s.Schedule(
		func() {
			log.Println("Every minute job")
		},
		scheduler.CronRule{
			Minute:  scheduler.EveryMinute,
			Hour:    scheduler.EveryHour,
			Day:     scheduler.EveryDay,
			Month:   scheduler.EveryMonth,
			Weekday: scheduler.EveryWeekday,
		},
	)
	log.Printf("Added first job with id %d", j1)

	// Schedule a job to run once (1min after execution)
	nextMoment := time.Now().Add(1 * time.Minute)
	j2 := s.Schedule(
		func() {
			log.Println("Single execution job")
		},
		scheduler.CronRule{
			Minute:  nextMoment.Minute(),
			Hour:    nextMoment.Hour(),
			Day:     nextMoment.Day(),
			Month:   nextMoment.Month(),
			Weekday: nextMoment.Weekday(),
		},
	)
	log.Printf("Added second job with id %d", j2)

	// Running scheduler
	s.Run()
	for {

	}
}
