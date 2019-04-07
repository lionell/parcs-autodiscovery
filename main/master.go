package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("udp", "255.255.255.255:1234")
	defer conn.Close()

	for i := 0; i < 10; i++ {
		conn.Write([]byte(fmt.Sprintf("Test %d", i)))
		time.Sleep(1 * time.Second)
	}
}
