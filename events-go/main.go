package main

import (
	"fmt"
	"time"

	"github.com/kelindar/event"
)

const (
	UserCreatedType uint32 = 1
	OrderPlacedType uint32 = 2
)

type UserCreated struct {
	ID   int
	Name string
}

func (e UserCreated) Type() uint32 {
	return UserCreatedType
}

type OrderPlaced struct {
	OrderID int
	UserID  int
	Amount  float64
}

func (e OrderPlaced) Type() uint32 {
	return OrderPlacedType
}

func main() {
	bus := event.NewDispatcher()

	defer event.Subscribe(bus, func(e UserCreated) {
		fmt.Printf("User created: ID=%d, Name=%s\n", e.ID, e.Name)
	})()

	defer event.Subscribe(bus, func(e OrderPlaced) {
		fmt.Printf("Order placed: OrderID=%d, UserID=%d, Amount=%.2f\n", e.OrderID, e.UserID, e.Amount)
	})()

	defer event.Subscribe(bus, func(e UserCreated) {
		fmt.Printf("Sending welcome message to user %s\n", e.Name)
	})()

	event.Publish(bus, UserCreated{ID: 1, Name: "Alice"})
	event.Publish(bus, UserCreated{ID: 2, Name: "Bob"})

	time.Sleep(100 * time.Millisecond)

	event.Publish(bus, OrderPlaced{OrderID: 101, UserID: 1, Amount: 99.99})
	event.Publish(bus, OrderPlaced{OrderID: 102, UserID: 2, Amount: 149.50})

	time.Sleep(100 * time.Millisecond)
}
