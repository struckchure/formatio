package lib

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	c *cron.Cron
}

type ScheduleArgs struct {
	Frequency string
	Duration  time.Duration
	Callback  func() error
}

type ScheduleOption func(*ScheduleArgs)

func WithFrequency(frequency string) ScheduleOption {
	return func(s *ScheduleArgs) {
		s.Frequency = frequency
	}
}

func WithDuration(duration time.Duration) ScheduleOption {
	return func(s *ScheduleArgs) {
		s.Duration = duration
	}
}

func WithCallback(callback func() error) ScheduleOption {
	return func(s *ScheduleArgs) {
		s.Callback = callback
	}
}

func (s *Scheduler) Schedule(opts ...ScheduleOption) error {
	args := &ScheduleArgs{}

	for _, opt := range opts {
		opt(args)
	}

	spec := fmt.Sprintf("%s %fh", args.Frequency, args.Duration.Hours())
	s.c.AddFunc(
		spec,
		func() {
			args.Callback()
		},
	)

	log.Println("Cron scheduled for: ", spec)

	return nil
}

func (s *Scheduler) Start() error {
	s.c.Start()

	return nil
}

func (s *Scheduler) Stop() error {
	s.c.Stop()

	return nil
}

func NewScheduler(c *cron.Cron) *Scheduler {
	return &Scheduler{c: c}
}

func NewSchedulerCron() *cron.Cron {
	c := cron.New()

	return c
}
