package parcs

import (
	"context"
	"log"
	"net"
	"time"
)

const port = ":4321"
const helloWorld = "Hello, world!"

// TODO(xlionell): Take context.Context as a parameter and stop on <-ctx.Done()
func DiscoverMaster() (net.IP, error) {
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Printf("couldn't resolve address to listen to: %v", err)
		return nil, err
	}
	conn, err := net.ListenUDP("udp", addr)
	defer conn.Close()
	if err != nil {
		log.Printf("couldn't listen for the connection: %v", err)
		return nil, err
	}

	buf := make([]byte, 1024)
	for {
		len, from, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("couldn't read the message from the master: %v", err)
			return nil, err
		}
		if string(buf[:len]) == helloWorld {
			return from.IP, nil
		}
	}
}

func Broadcast(ctx context.Context) {
	go cast(ctx, "255.255.255.255")
	go cast(ctx, "127.0.0.1")
}

func cast(ctx context.Context, address string) {
	conn, err := net.Dial("udp", address+port)
	if err != nil {
		log.Fatalf("While dialing %v:%v got error: %v", address, port, err)
	}
	defer conn.Close()

	log.Printf("Casting to %v", address)
	for {
		conn.Write([]byte(helloWorld))
		select {
		case <-ctx.Done():
			log.Printf("Finished casting to %v", address)
			return
		case <-time.After(1 * time.Second):
		}
	}
}
