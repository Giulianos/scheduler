package scheduler

import (
	"time"
)

// Scheduler contains a list of
// scheduled jobs
type Scheduler struct {
	jobs    []scheduledJob
	timer   *time.Timer
	running bool
}

// Job is the unit of work
// that can be scheduled
type Job func()

const (
	EveryMinute  int          = -1
	EveryHour    int          = -1
	EveryDay     int          = -1
	EveryMonth   time.Month   = -1
	EveryWeekday time.Weekday = -1
)

// CronRule is a rule
// that defines when a job is executed
type CronRule struct {
	Minute  int
	Hour    int
	Day     int
	Month   time.Month
	Weekday time.Weekday
}

type scheduledJob struct {
	job  Job
	rule CronRule
}

// New creates a new scheduler
func New() Scheduler {
	return Scheduler{
		jobs: []scheduledJob{},
	}
}

// Schedule schedules j for execution
func (s *Scheduler) Schedule(j Job, r CronRule) {
	s.jobs = append(s.jobs, scheduledJob{
		job:  j,
		rule: r,
	})
}

// Run starts the scheduler
func (s *Scheduler) Run() {
	s.scheduleNextExec(time.Now())
}

// Stop stops the execution of the
// scheduler
func (s *Scheduler) Stop() {
	s.timer.Stop()
}

func (s Scheduler) nextJobTime() time.Time {
	return time.Now().Add(time.Minute).Truncate(time.Minute)
}

func (r CronRule) isNow() bool {
	n := time.Now()
	if r.Minute >= 0 && r.Minute != n.Minute() {
		return false
	}
	if r.Hour >= 0 && r.Hour != n.Hour() {
		return false
	}
	if r.Day >= 0 && r.Day != n.Day() {
		return false
	}
	if r.Month >= 0 && r.Month != n.Month() {
		return false
	}
	if r.Weekday >= 0 && r.Weekday != n.Weekday() {
		return false
	}
	return true
}

func (s Scheduler) runJobsAsync() {
	for _, j := range s.jobs {
		if j.rule.isNow() {
			go j.job()
		}
	}
}

func (s *Scheduler) scheduleNextExec(t time.Time) {
	// Set timer for next execution
	s.timer = time.AfterFunc(
		time.Until(t),
		func() {
			s.runJobsAsync()
			s.scheduleNextExec(s.nextJobTime())
		},
	)
}
