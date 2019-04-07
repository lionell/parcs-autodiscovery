package main

import (
	"context"
	"log"
	"time"

	"github.com/lionell/parcs"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	errCh := parcs.Broadcast(ctx)
	for {
		select {
		case err := <-errCh:
			log.Printf("error while broadcasting: %v", err)
		case <-time.After(20 * time.Second):
			cancel()
			return
		}
	}
}
