package main_test

import (
	"github.com/s1ntaxe770r/lamport/pkg/clock"
	"github.com/s1ntaxe770r/lamport/pkg/node"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSingleTick(t *testing.T) {
	messageChan := make(chan node.Message, 100)
	OrderService := node.NewService("OrderService", messageChan)

	OrderService.Receive(clock.Received, OrderService.Clock.CurrentTimestamp())

	assert.Equal(t, 1, int(OrderService.Clock.CurrentTimestamp()))

}

func TestSend(t *testing.T) {
	messageChan := make(chan node.Message, 100)

	OrderService := node.NewService("OrderService", messageChan)
	go OrderService.HandleMessages()

	PaymentService := node.NewService("PaymentService", messageChan)
	CurrentTimestamp := PaymentService.Clock.CurrentTimestamp()
	PaymentService.Send(CurrentTimestamp)

	// Give some time for the message to be processed
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 2, int(OrderService.Clock.CurrentTimestamp()))
}
