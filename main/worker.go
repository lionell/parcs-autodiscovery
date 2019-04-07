package main

import (
	"flag"
	"log"
	"net"

	"github.com/lionell/parcs"
)

var masterIP net.IP

func main() {
	var masterIPFlag = flag.String("master", "", "IP of the master server.")
	flag.Parse()

	if *masterIPFlag != "" {
		masterIP = net.ParseIP(*masterIPFlag)
	} else {
		log.Println("Trying to discover the master IP address...")
		var err error
		masterIP, err = parcs.DiscoverMaster()
		if err != nil {
			log.Fatalf("couldn't discover the master: %v", err)
		}
	}
	log.Printf("Master IP address: %v.\n", masterIP)
}
