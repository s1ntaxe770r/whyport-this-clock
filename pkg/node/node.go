package node

import (
	"log/slog"

	"github.com/s1ntaxe770r/lamport/pkg/clock"
)

// ensure the service  struct implements the Node interface
var _ Node = &Service{}

// Node represents a service or individual process
type Node interface {

	// Send a message and return the resulting timestamp
	Send(currentClock string) int32
	// Receive a message
	Receive(event clock.Event, timestamp int32)

	// return the id of the node
	ID() string
}

type Service struct {
	Name   string
	Id     string
	logger *slog.Logger
	// TODO: all services should have a shared channel for message passing

	// Each Service gets a lamport clock
	clock clock.LamportClock
}

func (s *Service) Send(currentClock string) int32 {
	// Increment counter
	s.clock.Tick(s.clock.CurrentTimestamp())

	//TODO: Send message over channel

	return s.clock.CurrentTimestamp()
}

func (s *Service) Receive(event clock.Event, timestamp int32) {
	s.clock.Tick(timestamp)
}

func (s *Service) ID() string {

	return s.Id

}
