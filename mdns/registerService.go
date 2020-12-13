package mdns

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/grandcat/zeroconf"
)

// RegisterService for registration of _yottamusic service over mDNS
func RegisterService() {

	name := "YottaMusic"          //flag.String("name", "YottaMusic", "Name for the Service.")
	service := "_yottamusic._tcp" //flag.String("service", "_yottamusic._tcp", "Set the service type of the new Service.")
	domain := "local."            //flag.String("domain", "local.", "Set the network Domain.")
	port := 80                    //flag.Int("port", 80, "Set the port the service is listening to.")

	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	} else {
		// fmt.Printf("Device Name: %s", hostName)
		name = string(hostName)
	}

	// Metadata information about the service
	metadata := []string{
		"version = 0.1",
		"developer = Shachindra",
	}

	mDNSService, err := zeroconf.Register(name, service, domain, port, metadata, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer mDNSService.Shutdown()
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
