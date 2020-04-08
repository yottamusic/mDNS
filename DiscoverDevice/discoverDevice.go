package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/grandcat/zeroconf"
)

var (
	service  = flag.String("service", "_yottamusic._tcp", "Set the Service Category to look for Devices.")
	domain   = flag.String("domain", "local", "Set the search Domain.")
	waitTime = flag.Int("wait", 10, "Duration in [s] to run Discovery.")
)

func serviceCall(ip string, port int) {
	url := fmt.Sprintf("http://%v:%v", ip, port)

	log.Println("Making GET Request to", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Response: %s\n", data)
}

func main() {

	flag.Parse()

	// Discover all Services on the Network (For _yottamusic._tcp)
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			log.Println(entry)
			log.Println("Found service:", entry.ServiceInstanceName(), entry.Text)
			serviceCall(entry.AddrIPv4[0].String(), entry.Port)
		}
		log.Println("No more entries.")
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*waitTime))
	defer cancel()
	err = resolver.Browse(ctx, *service, *domain, entries)
	if err != nil {
		log.Fatalln("Failed to Browse:", err.Error())
	}

	<-ctx.Done()

	// Wait some additional time to see debug messages on go routine shutdown.
	time.Sleep(1 * time.Second)
}
