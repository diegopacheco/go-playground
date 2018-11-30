package main

import (
	"context"
	"time"

	"github.com/signalfx/golib/datapoint"
	"github.com/signalfx/golib/event"
)

func main() {

	client := NewHTTPSink()
	client.AuthToken = "ABCD..."
	ctx := context.Background()
	client.AddDatapoints(ctx, []*datapoint.Datapoint{
		GaugeF("hello.world", nil, 1.0),
	})
	dims := make(map[string]string)
	client.AddEvents(ctx, []*event.Event{
		event.New("hello.world", event.USERDEFINED, dims, time.Time{}),
	})

}
