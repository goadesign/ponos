# Replicated Map

Replicated maps provide a mechanism for sharing data across a fleet of
microservices and receiving events when the data changes.

## Overview

Ponos replicated maps leverage Redis hashes and pub/sub to maintain replicated
in-memory copies of a map across multiple nodes. Any change to the map is
automatically replicated to all nodes and results in a notification that can be
used to trigger actions.

Upon joining a replicated map the node receives an up-to-date snapshot of its
content followed by updates as they occur.

## Usage

### Creating a Replicated Map

To create a replicated map you must provide a name and a Redis client. The name
is used to namespace the Redis keys and pub/sub channels used by the map.

```go
package main

import (
        "context"

        "github.com/redis/go-redis/v9"
        "goa.design/ponos/rmap"
)

func main() {
        // Create a Redis client
        client := redis.NewClient(&redis.Options{
            Addr: "localhost:6379",
        })
    
        // Join or create a replicated map
        m, err := rmap.Join(context.Background(), "my-map", client)
        if err != nil {
            panic(err)
        }
}
```

### Reading and Writing to the Map

The `Map` method returns a copy of the current map. The `Get` method returns the
value for a given key. The `Set` method sets the value for a given key. 

```go
        // Print the current contents of the map
        fmt.Println(m.Map())

        // Add a new key
        m.Set(context.Background(), "foo", "bar")

        // Keys set by the current process are available immediately
        val, ok := m.Get("foo") // ok is true
```

### Cleaning Up

When you are done with a replicated map you should cancel the context used to
create it. This will cause the map to unsubscribe from the pub/sub channel and
stop receiving updates.

```go
        ctx, cancel := context.WithCancel(context.Background())
        // Cleanup when done
        defer cancel()

        // Join or create a replicated map
        m, err := rmap.Join(ctx, "my-map", client)
        if err != nil {
            panic(err)
        }
```

## Example

The following example creates a replicated map and then starts a goroutine that
listens to notifications. The main goroutine then sets 100 keys and waits for
the notifications to be received.

```go
package main

import (
        "context"
        "fmt"
        "strconv"
        "sync"

        "github.com/redis/go-redis/v9"
        "goa.design/ponos/rmap"
)

func main() {
        // Create a Redis client
        client := redis.NewClient(&redis.Options{
                Addr: "localhost:6379",
                Password: os.Getenv("REDIS_PASSWORD"),
        })
    
        ctx, cancel := context.WithCancel(context.Background())
        // Cleanup when done
        defer cancel()

        // Join or create a replicated map
        m, err := rmap.Join(ctx, "my-map", client)
        if err != nil {
                panic(err)
        }

        // Start a goroutine to listen for updates
        var wg sync.WaitGroup
        wg.Add(1)
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        case <-m.C:
                                if len(m.Map()) == 100 {
                                        // We received all the updates
                                        wg.Done()
                                }
                        }
                }
        }()

        // Send a few updates
        for i := 0; i < 100; i++ {
                m.Set(ctx, "foo-" + strconv.Itoa(i+1), "bar")
        }

        // Wait for the updates to be received
        wg.Wait()
}
```