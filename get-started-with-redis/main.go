package main

import (
	"fmt"
	"log"
	"os"
)

const (
	dbconn = "localhost:6379"
	db     = 0
)

func usage() {
	fmt.Println(`Usage: go run main.go <demo>
Where <demo> is one of:
	- ping: run the ping command
	- getset: get and set a string value
	- expire: set an expiring key
	- pipeline: run a batch of commands
	- transaction: run a batch of commands that must all succeed
	- pubsub: send a message to a channel and listen to the channel
	- reset: restore the initial data set`)
}

func run() error {
	// Create a new client from a connection string and a database number (0-15)
	client := newClient(dbconn, 0)

	if len(os.Args) < 2 {
		usage()
		return nil
	}

	switch os.Args[1] {
	case "ping":
		// Ping the redis server and fetch some database information.
		fmt.Printf("\nPing: Test the connection\n")
		ping(client)

	case "getset":
		// Get and set a string value.
		fmt.Printf("\nGet/Set: Update the quest to title case\n")
		err := getAndSet(client)
		if err != nil {
			return fmt.Errorf("getAndSet failed: %w", err)
		}

	case "expire":
		// Set an expiring key and wait for it to expire
		fmt.Printf("\nExpire: Add a player temporarily\n")
		err := expiringKeys(client)
		if err != nil {
			return fmt.Errorf("expiringKeys failed: %w", err)
		}

	case "pipeline":
		// Run a batch of commands.
		fmt.Printf("\nPipeline: Update score and challenges_completed for team Snarkdumbthimble\n")
		err := pipeline(client)
		if err != nil {
			return fmt.Errorf("pipeline failed: %w", err)
		}

	case "transaction":
		// Run several commands that must all succeed.
		// If any of them fails, the transaction will be canceled.
		fmt.Printf("\nTransaction: Rearrange the teams\n")
		err := transaction(client)
		if err != nil {
			return fmt.Errorf("transaction failed: %w", err)
		}

	case "pubsub":
		// Send messages to a publish/receive channel.
		// Listen to the channel and receive the messages.
		fmt.Printf("\nPub/Sub: Publish challenges to subscribed teams\n")
		err := pubsub(client)
		if err != nil {
			return fmt.Errorf("pubsub stopped: %w", err)
		}

	case "reset":
		// Reset the database
		fmt.Printf("\nReset the database\n")
		err := resetdata(client)
		if err != nil {
			return fmt.Errorf("cannot reset database: %w", err)
		}
	default:
		usage()
	}
	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
