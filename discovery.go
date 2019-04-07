package parcs

import (
	"context"
	"log"
	"net"
	"time"
)

const port = ":4321"
const helloWorld = "Hello, world!"

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

func Broadcast(ctx context.Context) error {
	conn, err := net.Dial("udp", "255.255.255.255"+port)
	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		log.Print("Broadcasting 'hello world' message.")
		_, err := conn.Write([]byte(helloWorld))
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			log.Print("Finished broadcasting.")
			return nil
		case <-time.After(1 * time.Second):
		}
	}
}
