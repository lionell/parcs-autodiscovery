package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/lionell/parcs"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	parcs.Broadcast(ctx)
	ch := parcs.Listen()
	ipList := make([]net.IP, 0)
outer:
	for {
		select {
		case ip := <-ch:
			ipList = append(ipList, ip)
		case <-time.After(20 * time.Second):
			break outer
		}
	}
	cancel()
	f, err := os.Create("hosts.lst")
	if err != nil {
		log.Fatalf("can't open file for writing: %v", err)
		return
	}
	defer f.Close()
	for _, ip := range ipList {
		f.WriteString(fmt.Sprintf("%v\n", ip))
	}
}
