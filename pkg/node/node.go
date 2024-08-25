package node

import (
	"github.com/s1ntaxe770r/lamport/pkg/clock"
	"github.com/teris-io/shortid"
	"log/slog"
	"os"
)

// ensure the service  struct implements the Node interface
var _ Node = &Service{}

// Node represents a service or individual process
type Node interface {

	// Send a message and return the resulting timestamp
	Send(CurrentTimestamp int32) int32
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

	// Each Service gets a lamport Clock
	Clock clock.LamportClock
}

func NewService(name string) *Service {
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	return &Service{
		Name:   name,
		Id:     sid.String(),
		logger: logger,
		Clock:  clock.LamportClock{},
	}

}

func (s *Service) Send(CurrentTimestamp int32) int32 {
	// Increment counter
	s.Clock.Local()
	//TODO: Send message over channel
	return s.Clock.CurrentTimestamp()
}

func (s *Service) Receive(event clock.Event, timestamp int32) {
	s.Clock.Tick(timestamp)
}

func (s *Service) ID() string {
	return s.Id

}
