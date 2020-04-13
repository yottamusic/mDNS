package registerDevice

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/grandcat/zeroconf"
)

func RegisterDevice() {

	name := "YottaMusic"          //flag.String("name", "YottaMusic", "Name for the Service.")
	service := "_yottamusic._tcp" //flag.String("service", "_yottamusic._tcp", "Set the service type of the new Service.")
	domain := "local."            //flag.String("domain", "local.", "Set the network Domain.")
	port := 80                    //flag.Int("port", 80, "Set the port the service is listening to.")

	hostName, err := ioutil.ReadFile("/etc/hostname")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Device Name: %s", hostName)
		name = string(hostName)
	}

	server, err := zeroconf.Register(name, service, domain, port, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()
	log.Println("Published service:")
	log.Println("- Name:", name)
	log.Println("- Type:", service)
	log.Println("- Domain:", domain)
	log.Println("- Port:", port)

	// Clean exit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sig:
		// Exit by user
		// case <-tc:
		// Exit by timeout
	}

	log.Println("Shutting down mDNS Registration")
	return
}
