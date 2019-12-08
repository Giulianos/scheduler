package scheduler

import (
	"sync"
	"time"
)

// Scheduler contains a list of
// scheduled jobs
type Scheduler struct {
	timer   *time.Timer
	running bool

	jobsMu sync.Mutex
	jobs   map[ScheduledJobID]scheduledJob
	nextID ScheduledJobID
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
	Weekday time.Weekday
	Month   time.Month
	Day     int
	Hour    int
	Minute  int
}

// ScheduledJobID is an ID provided to each
// job that is scheduled, it is used to remove
// it from the scheduler if needed
type ScheduledJobID int64

type scheduledJob struct {
	job  Job
	rule CronRule
}

// New creates a new scheduler
func New() Scheduler {
	return Scheduler{
		jobs: map[ScheduledJobID]scheduledJob{},
	}
}

// Schedule schedules j for execution
func (s *Scheduler) Schedule(j Job, r CronRule) ScheduledJobID {
	s.jobsMu.Lock()
	jobID := s.nextID
	s.jobs[jobID] = scheduledJob{
		job:  j,
		rule: r,
	}
	s.nextID++
	s.jobsMu.Unlock()
	return jobID
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

// RemoveJob removes a job from the list of scheduled
// jobs
func (s *Scheduler) RemoveJob(id ScheduledJobID) {
	s.jobsMu.Lock()
	delete(s.jobs, id)
	s.jobsMu.Unlock()
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
