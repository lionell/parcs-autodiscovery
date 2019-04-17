package parcs

import (
	"log"
	"net"
	"strings"
)

const RegistrationPort = ":1234"

func Register(address net.IP) error {
	conn, err := net.Dial("tcp", address.String()+RegistrationPort)
	if err != nil {
		log.Printf("couldn't connect to master address: %v", err)
		return err
	}
	log.Printf("Connected to: %v", address)
	defer conn.Close()
	return nil
}

func extractIP(address net.Addr) net.IP {
	return net.ParseIP(strings.Split(address.String(), ":")[0])
}

func Listen() <-chan net.IP {
	ch := make(chan net.IP)
	go listen(ch)
	return ch
}

func listen(ch chan<- net.IP) {
	l, err := net.Listen("tcp", RegistrationPort)
	if err != nil {
		log.Printf("couldn't listen on port %v: %v", RegistrationPort, err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("couldn't accept connection: %v", err)
		}
		ch <- extractIP(conn.RemoteAddr())
		conn.Close()
	}
}
