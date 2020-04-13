package main

import (
	"flag"
	"fmt"

	"github.com/yottamusic/mDNS/discoverDevice"
	"github.com/yottamusic/mDNS/registerDevice"
)

func main() {
	fmt.Println("YottaMusic mDNS Device Registration and Discovery")

	var operation string
	flag.StringVar(&operation, "operation", "register", "Register or Discover YottaMusic Devices")
	flag.Parse()
	fmt.Println("Starting to ", operation)

	if operation == "register" {
		registerDevice.RegisterDevice()
	} else if operation == "discover" {
		discoverDevice.DiscoverDevice()
	} else {
		fmt.Println("Invalid Operation")
	}
}
