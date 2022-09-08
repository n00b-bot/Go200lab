package pubsub

import (
	"strconv"
	"time"
)

type Message struct {
	id        string
	channel   Topic
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	return &Message{
		id:        strconv.Itoa(int(time.Now().UTC().UnixNano())),
		data:      data,
		createdAt: time.Now(),
	}
}

func (m *Message) String() string {
	return string(m.channel)
}

func (m *Message) Channel() Topic {
	return m.channel
}

func (m *Message) SetChannel(channel Topic) {
	m.channel = channel
}

func (m *Message) Data() interface{} {
	return m.data
}
