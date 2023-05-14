package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"goa.design/ponos/streaming"
)

// NOTE: the example below does not handle errors for brevity.
func main() {
	// Create Redis client
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: os.Getenv("REDIS_PASSWORD")})
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	// Create stream
	stream, err := streaming.NewStream(ctx, "my-stream", rdb)
	if err != nil {
		panic(err)
	}

	// Don't forget to destroy the stream when done
	defer stream.Destroy(ctx)

	// Add a new event to topic "my-topic"
	id1, err := stream.Add(ctx, "event 1", []byte("payload 2"), streaming.WithTopic("my-topic"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("event 1 id: %s\n", id1)

	// Add a new event to topic "my-other-topic"
	id2, err := stream.Add(ctx, "event 2", []byte("payload 2"), streaming.WithTopic("my-other-topic"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("event 2 id: %s\n", id2)

	// Create sink for stream "my-stream" that reads from the beginning and
	// waits for events for up to 100ms
	sink, err := stream.NewSink(ctx, "my-sink",
		streaming.WithSinkStartAtOldest(),
		streaming.WithSinkBlockDuration(100*time.Millisecond))
	if err != nil {
		panic(err)
	}

	// Don't forget to close the sink when done
	defer sink.Close()

	// Read both events
	event := <-sink.C
	fmt.Printf("topic: %s, event: %s, payload: %s\n", event.Topic, event.EventName, event.Payload)
	if err := sink.Ack(ctx, event); err != nil {
		panic(err)
	}

	event = <-sink.C
	fmt.Printf("topic: %s, event: %s, payload: %s\n", event.Topic, event.EventName, event.Payload)
	if err := sink.Ack(ctx, event); err != nil {
		panic(err)
	}
}