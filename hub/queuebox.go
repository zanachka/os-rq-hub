package hub

import (
	"sync"

	"github.com/cfhamlet/os-rq-pod/pod"
)

// UpstreamMap TODO
type UpstreamMap map[UpstreamID]*Upstream

// QueueUpstreamsMap TODO
type QueueUpstreamsMap map[pod.QueueID]QueueUpstreams

// QueueUpstreams TODO
type QueueUpstreams struct {
	queue   *Queue
	streams UpstreamMap
}

// QueueBox TODO
type QueueBox struct {
	hub               *Hub
	queueUpstreamsMap QueueUpstreamsMap
	locker            *sync.RWMutex
}

// NewQueueBox TODO
func NewQueueBox(hub *Hub) *QueueBox {
	return &QueueBox{
		hub,
		QueueUpstreamsMap{},
		&sync.RWMutex{},
	}
}