package main

import (
	"context"
	"time"

	"github.com/lionell/parcs"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	parcs.Broadcast(ctx)
	time.Sleep(20 * time.Second)
	cancel()
}
