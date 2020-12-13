package main

import (
	"flag"
	"fmt"

	"github.com/yottamusic/mDNS/mdns"
)

func main() {
	fmt.Println("YottaMusic mDNS Device Registration and Discovery")

	var operation string
	flag.StringVar(&operation, "operation", "register", "Register or Discover YottaMusic Devices")
	flag.Parse()
	fmt.Println("Starting to", operation)

	if operation == "register" {
		mdns.RegisterService()
	} else if operation == "discover" {
		mdns.DiscoverService()
	} else {
		fmt.Println("Invalid Operation")
	}
}
