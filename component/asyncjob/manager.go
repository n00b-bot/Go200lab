package asyncjob

import (
	"context"
)

type manager struct {
	jobs         []Job
	isConcurrent bool
}

func NewManager(c bool, jobs ...Job) *manager {
	return &manager{
		jobs:         jobs,
		isConcurrent: c,
	}
}

func (m *manager) Run(ctx context.Context) error {
	errs := make(chan error, len(m.jobs))
	//m.wg.Add(len(m.jobs))
	for i, _ := range m.jobs {
		if m.isConcurrent {
			go func(j Job) {
				errs <- m.runJob(ctx, j)
				//m.wg.Done()
			}(m.jobs[i])
			continue
		}
		errs <- m.runJob(ctx, m.jobs[i])
		//m.wg.Done()
	}

	//m.wg.Wait()
	var err error
	for i := 1; i <= len(m.jobs); i++ {
		if v := <-errs; v != nil {
			err = v
		}
	}
	return err
}

func (m *manager) runJob(ctx context.Context, job Job) error {
	if err := job.Execute(ctx); err != nil {
		for {
			if job.State() == RetryFailed {
				return err
			}
			if job.Retry(ctx) == nil {
				return nil
			}
		}
	}
	return nil
}
