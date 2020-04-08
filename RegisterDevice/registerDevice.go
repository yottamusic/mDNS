package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/grandcat/zeroconf"
)

var (
	name    = flag.String("name", "YottaMusic", "Name for the Service.")
	service = flag.String("service", "_yottamusic._tcp", "Set the service type of the new Service.")
	domain  = flag.String("domain", "local.", "Set the network Domain.")
	port    = flag.Int("port", 80, "Set the port the service is listening to.")
)

func main() {

	flag.Parse()
	hostName, err := ioutil.ReadFile("/etc/hostname")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Device Name: %s", hostName)
		*name = string(hostName)
	}

	server, err := zeroconf.Register(*name, *service, *domain, *port, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()
	log.Println("Published service:")
	log.Println("- Name:", *name)
	log.Println("- Type:", *service)
	log.Println("- Domain:", *domain)
	log.Println("- Port:", *port)

	// Clean exit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sig:
		// Exit by user
		// case <-tc:
		// Exit by timeout
	}

	log.Println("Shutting down.")
}
