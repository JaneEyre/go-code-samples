package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Goal: Read multiple files concurrently. All files must be read successfully.
// If one goroutine fails, cancel all the other goroutines.
//
// The cancelled goroutines can inspect the error from canceling the context with ctx.Err()..

func ReadFilesConcurrently(paths []string) ([][]byte, error) {
	var contents [][]byte

	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	resCh := make(chan []byte)
	wg := sync.WaitGroup{}

	for _, path := range paths {
		wg.Add(1)
		go func(ctx context.Context, cancel context.CancelCauseFunc, p string, resCh chan<- []byte, wg *sync.WaitGroup) {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond) // simulate workload
			select {
			case <-ctx.Done():
				log.Printf("ReadFilesConcurrently (goroutine): Context canceled for path %s: %v", p, ctx.Err())
				wg.Done()
				return
			default:
				// If an error occurs here, cancel all the other goroutines
				content, err := ReadFile(p)
				if err != nil {
					cancel(fmt.Errorf("error reading %s: %w", p, err))
					log.Printf("ReadFilesConcurrently (goroutine): Context canceled for path %s: %v", p, err)
					wg.Done()
					return
				}
				resCh <- content
			}
		}(ctx, cancel, path, resCh, &wg)
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	for c := range resCh {
		contents = append(contents, c)
	}

	if e := ctx.Err(); e != nil {
		return nil, fmt.Errorf("ReadFilesConcurrently: %w", e)
	}
	return contents, nil
}
