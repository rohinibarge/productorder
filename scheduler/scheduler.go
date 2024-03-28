package scheduler

import (
	"fmt"
	"time"
)

type Task interface {
	Execute()
}

type Cron struct {
	Interval time.Duration
	Task     Task
}

func NewScheduler(interval time.Duration, task Task) *Cron {
	return &Cron{
		Interval: interval,
		Task:     task,
	}
}

func (s *Cron) Start() {
	ticks := time.Tick(s.Interval)
	for now := range ticks {
		s.Task.Execute()
		fmt.Println("task triggered at ", now)
	}
}
