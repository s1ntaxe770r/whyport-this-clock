package clock

import (
	"sync"
	"sync/atomic"
)

type Event int

const (
	Send = iota
	Received
	Local
)

type LamportClock struct {
	counter atomic.Int32
	mutext  sync.Mutex
}

// Tick when an Event occurs
func (lc *LamportClock) Tick(currentCLock int32) {
	currentTime := max(lc.counter.Load(), currentCLock) + 1
	lc.counter.Add(currentTime)
}

func (lc *LamportClock) Local() {
	lc.counter.Add(1)
}

func (lc *LamportClock) CurrentTimestamp() int32 {
	return lc.counter.Load()
}
