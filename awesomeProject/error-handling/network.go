package main

import (
	"errors"
	"fmt"
	"log"
	"net"
)

// Connect to a TCP server and check the error. use errors.As() to unwrap the net.OpError, and test if the error is transient.
func connectToTCPServer() error {
	var err error
	var conn net.Conn
	for retry := 3; retry > 0; retry-- {
		conn, err = net.Dial("tcp", "127.0.0.1:12345")
		if err != nil {
			// Check if err is a net.OpError
			opErr := &net.OpError{}
			if errors.As(err, &opErr) {
				log.Println("err is net.OpError:", opErr.Error())
				// test if the error is temporary
				if opErr.Temporary() {
					log.Printf("Retrying...\n")
					continue
				}
				retry = 0
			}
		}
	}
	if err != nil {
		return fmt.Errorf("connect failed: %w", err)
	}
	defer conn.Close()
	// send or receive data
	return nil
}
