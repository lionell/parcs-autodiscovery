package main

import (
	"context"
	"time"

	"github.com/lionell/parcs"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	go parcs.Broadcast(ctx)
	time.Sleep(20 * time.Second)
}
