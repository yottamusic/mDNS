package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("YottaMusic mDNS Device Registration and Discovery")

	var operation string
	flag.StringVar(&operation, "operation", "register", "Register or Discover YottaMusic Devices")
	flag.Parse()

	if operation == "register" {
		registerDevice.registerDevice()
	} else if operation == "discover" {
		discoverDevice.discoverDevice()
	} else {
		fmt.Println("Invalid Operation")
	}
}
