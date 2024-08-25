package node

import (
	"github.com/s1ntaxe770r/lamport/pkg/clock"
	"github.com/teris-io/shortid"
	"log/slog"
	"os"
	"sync"
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

type Message struct {
	SenderID  string
	Timestamp int32
	Content   string
}

type Service struct {
	Name        string
	Id          string
	logger      *slog.Logger
	MessageChan chan Message
	// Each Service gets a lamport Clock
	Clock  clock.LamportClock
	mutext sync.RWMutex
}

func NewService(name string, messageChan chan Message) *Service {
	sid, _ := shortid.Generate()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	return &Service{
		Name:        name,
		Id:          sid,
		logger:      logger,
		Clock:       clock.LamportClock{},
		MessageChan: messageChan,
	}

}

func (s *Service) Send(CurrentTimestamp int32) int32 {
	// Increment counter
	s.Clock.Local()

	msg := Message{
		SenderID:  s.Id,
		Timestamp: s.Clock.CurrentTimestamp(),
		Content:   "HI LESLIE!!!",
	}

	s.mutext.Lock()
	s.MessageChan <- msg
	s.mutext.Unlock()

	s.logger.Info("sent message", "Timestamp", msg.Timestamp, "Content", msg.Content)
	return s.Clock.CurrentTimestamp()
}

func (s *Service) Receive(event clock.Event, timestamp int32) {
	s.Clock.Tick(timestamp)
}

func (s *Service) ID() string {
	return s.Id

}

func (s *Service) HandleMessages() {
	s.logger.Info("Starting message handler")
	for msg := range s.MessageChan {
		s.Receive(clock.Received, msg.Timestamp)
		s.logger.Info("Received Message", "from", msg.SenderID, "timestamp", msg.Timestamp, "Content", msg.Content)
	}

}
