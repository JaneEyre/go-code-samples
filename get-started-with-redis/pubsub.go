package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// Task: Send each challenge to the "challenge" channel.
// Each team is subscribed to the "challenge" channel and enters the challenge upon receiving.

const (
	// the name of the PubSub channel
	pubsubChan = "challenge"
)

// Res is the result of reading the pubsub channel
type Res struct {
	result string
	err    error
}

// Team manages a team's subscription
// Each team uses its own client to subscribe to the "challenge" channel
type Team struct {
	name    string
	client  *redis.Client
	channel *redis.PubSub
}

// getTeams scans the database for "team:*" keys
// and returns a slice of Team structs with names filled from the keys
func getTeams(client *redis.Client) []Team {
	ctx := context.Background()
	teams := make([]Team, 3)
	teamsets := make([]string, 0, 3)
	keys := make([]string, 0, 3)
	var cursor uint64
	for {
		// Scan returns a slice of matches. The count may or may not be reached
		// in the first call to Scan, so the code needs to call Scan in a loop and
		// append the found keys to the teamsets slice until the cursor "returns to 0".
		var err error
		keys, cursor, err = client.Scan(ctx, cursor, "team:*", 3).Result()
		if err != nil {
			break
		}
		teamsets = append(teamsets, keys...)
		if cursor == 0 {
			break
		}
	}
	// Lazily assume that the scan has returned 3 team sets
	for i := 0; i < 3; i++ {
		teams[i].name = teamsets[i]
		// each team uses its own client
		teams[i].client = newClient(dbconn, 0)
	}
	return teams
}

// subscribe subscribes to the "challenge" channel
// and waits for the subscription to be completed
func (team *Team) subscribe() error {
	ctx := context.Background()
	// Subscribe to the "challenge" channel
	pubSub := team.client.Subscribe(ctx, pubsubChan)

	// The first Subscribe() call creates the channel.
	// Until that point, any attempt to publish something fails.
	reply, err := pubSub.Receive(ctx)
	if err != nil {
		return fmt.Errorf("subscribing to channel '%s' failed: %w", pubsubChan, err)
	}
	// Expected response type is "*Subscription". Otherwise, something failed.
	switch reply.(type) {
	case *redis.Subscription:
		// Success!
	case *redis.Message:
		// The channel is already active and contains messages, hence also a success
	case *redis.Pong:
		// letL's call it a success
	default:
		return fmt.Errorf("subscribing to a channel failed: received a reply of type %T, expected: *redis.Subscription", reply)
	}

	team.channel = pubSub

	fmt.Printf("%s subscribed to channel '%s'\n", team.name, pubsubChan)
	return nil
}

// receive receives messages from the "challenge" channel.
// It starts a goroutine that reads from the pubsub channel until
// the channel is closed or the context is done.
func (team *Team) receive(ctx context.Context, resChan chan<- Res) {
	ch := team.channel.Channel()
	defer close(resChan)
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				// The pubsub channel has been closed
				return
			}
			resChan <- Res{fmt.Sprintf("%s received challenge '%s'", team.name, msg.Payload), nil}
		case <-ctx.Done():
			resChan <- Res{"", ctx.Err()}
			return
		}
	}
}

// publish publishes the challenge to the "challenge" channel
func publish(client *redis.Client, challenge string) error {
	ctx := context.Background()
	fmt.Printf("publishing challenge '%s'\n", challenge)
	return client.Publish(ctx, pubsubChan, challenge).Err()
}

// pubsub subscribes to the "challenge" channel, publishes the challenges,
// and receives the published messages.
func pubsub(client *redis.Client) (err error) {
	ctx := context.Background()

	// Step 1: subscribe each team
	teams := getTeams(client)
	for i := 0; i < 3; i++ {
		err = teams[i].subscribe()
		if err != nil {
			return fmt.Errorf("subscribing failed: %w", err)
		}
	}

	// Step 2: publish challenges
	// Read the challenges from the sorted set "challenges" and publish them
	for i := int64(0); i < 5; i++ {
		challenge := client.ZRange(ctx, "challenges", i, i).Val()[0]
		err = publish(client, challenge)
		if err != nil {
			return fmt.Errorf("cannot publish challenge %s: %w", challenge, err)
		}
	}
	// Close the channel after one second, to terminate the receive loops.
	time.AfterFunc(time.Second, func() {
		teams[0].channel.Close()
		fmt.Println(`PubSub channel "challenges" closed`)
	})

	// Step 3: receive published messages
	rch := make(chan Res)
	for i := 0; i < 3; i++ {
		go teams[i].receive(ctx, rch)
	}
	for msg := range rch {
		if msg.err != nil {
			return fmt.Errorf("cannot receive challenge: %w", msg.err)
		}
		fmt.Println(msg.result)
	}

	return nil
}
