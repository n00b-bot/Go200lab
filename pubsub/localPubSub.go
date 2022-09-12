package pubsub

import (
	"sync"

	"golang.org/x/net/context"
)

type localPubSub struct {
	messQueue  chan *Message
	mapChannel map[Topic][]chan *Message
	locker     *sync.RWMutex
}

func NewPubSub() *localPubSub {
	pb := &localPubSub{
		messQueue:  make(chan *Message, 1000),
		mapChannel: make(map[Topic][]chan *Message),
		locker:     new(sync.RWMutex),
	}
	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, topic Topic, data *Message) error {
	data.SetChannel(topic)
	go func() {
		ps.messQueue <- data
	}()

	ps.run()
	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic Topic) (ch <-chan *Message, close func()) {
	c := make(chan *Message)
	ps.locker.Lock()
	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *Message{c}
	}
	ps.locker.Unlock()
	return c, func() {
		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)
					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}
}

func (ps *localPubSub) run() {
	go func() {
		for {
			m := <-ps.messQueue
			if sub, ok := ps.mapChannel[m.Channel()]; ok {
				for i := range sub {
					go func(c chan *Message) {
						c <- m 
					}(sub[i])
				}
			}
		}
	}()
}
