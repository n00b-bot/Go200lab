package asyncjob

import (
	"context"
	"time"
)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(time []time.Duration)
}

type JobState int

const (
	Init JobState = iota
	Running
	Failed
	Timeout
	Completed
	RetryFailed
)

var defaultRetryTime = []time.Duration{time.Second, time.Second * 5, time.Second * 10}

type JobHandler func(context.Context) error

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	return &job{
		config:     jobConfig{MaxTimeout: time.Second * 10, Retries: defaultRetryTime},
		handler:    handler,
		state:      Init,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}
}

func (j *job) Execute(ctx context.Context) error {
	j.state = Running
	if err := j.handler(ctx); err != nil {
		j.state = Failed
		return err
	}
	j.state = Completed
	return nil

}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])
	err := j.Execute(ctx)
	if err == nil {
		j.state = Completed
		return nil
	}
	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = RetryFailed
		return err
	}
	j.state = Failed
	return err
}

func (j *job) State() JobState { return j.state }

func (j *job) RetryIndex() int { return j.retryIndex }

func (j *job) SetRetryDurations(time []time.Duration) {
	if len(time) == 0 {
		return
	}
	j.config.Retries = time
}
