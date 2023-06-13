package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/redis/go-redis/v9"
	"io"
	"os"
)

const (
	setupfile = "testdata/setup.redis"
)

// resetdata flushes all data from the current database (db 0)
// and runs the commands from file setup.redis to set up the database.
func resetdata(client *redis.Client) error {
	ctx := context.Background()

	// read file "setup.redis" line by line
	setup, err := os.Open(setupfile)
	if err != nil {
		return fmt.Errorf("cannot open %s: %w", setupfile, err)
	}
	defer setup.Close()

	client.FlushDB(ctx) // FlushDB never fails.

	csv := csv.NewReader(setup)
	csv.Comma = ' '
	csv.FieldsPerRecord = -1 // Variable number of fields per line

	for {
		cmd, err := csv.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("csv: cannot read a line from %s: %w", setupfile, err)
		}

		// cmd is a slice of strings, Do() expects a slice of 'any'.
		// The memory layout of the two slice types is not the same,
		// so we need to convert cmd to a slice of 'any'.
		doCmd := make([]interface{}, len(cmd))
		for i, v := range cmd {
			doCmd[i] = v
		}

		err = client.Do(ctx, doCmd...).Err()
		if err != nil {
			return fmt.Errorf("resetdata: cannot execute '%v': %w", cmd, err)
		}
	}
	return nil
}
