package subscriber

import (
	"context"
	"food/common"
	"food/component/appctx"
	"food/component/asyncjob"
	"food/pubsub"
	"log"
)

type consumerJob struct {
	Title   string
	Handler func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appctx appctx.AppContext
}

func NewEngine(appCtx appctx.AppContext) *consumerEngine {
	return &consumerEngine{
		appctx: appCtx,
	}
}

func (e *consumerEngine) Start() error {
	e.startSubTopic(pubsub.Topic(common.UserLike), true, UpLike(e.appctx))
	return nil
}

func (e *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consconsumerJobs ...consumerJob) error {
	c, _ := e.appctx.GetPubSub().Subscribe(context.Background(), topic)
	getJobHandler := func(job *consumerJob, data *pubsub.Message) func(context.Context) error {
		return func(ctx context.Context) error {
			return job.Handler(ctx, data)
		}
	}

	go func() {
		msg := <-c
		jobs := make([]asyncjob.Job, len(consconsumerJobs))
		for i, v := range consconsumerJobs {
			job := getJobHandler(&v, msg)
			jobs[i] = asyncjob.NewJob(job)
		}
		manager := asyncjob.NewManager(isConcurrent, jobs...)
		if err := manager.Run(context.Background()); err != nil {
			log.Println(err)
		}
	}()
	return nil
}
