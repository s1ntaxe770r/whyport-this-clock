package main_test

import (
	"testing"

	"github.com/s1ntaxe770r/lamport/pkg/clock"
	"github.com/s1ntaxe770r/lamport/pkg/node"
	"github.com/stretchr/testify/assert"
)

func TestSingleTick(t *testing.T) {

	OrderService := node.NewService("OrderService")

	OrderService.Receive(clock.Received, OrderService.Clock.CurrentTimestamp())

	assert.Equal(t, 1, int(OrderService.Clock.CurrentTimestamp()))

}

func TestSend(t *testing.T) {

	PaymentService := node.NewService("PaymentService")

	CurrentTimestamp := PaymentService.Send(PaymentService.Clock.CurrentTimestamp())
	PaymentService.Receive(clock.Received, CurrentTimestamp)

	assert.Equal(t, 2, int(PaymentService.Clock.CurrentTimestamp()))
}
