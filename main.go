package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {

	port := flag.Int("p", 0, "Port number to use, default is 0 for random")
	dir := flag.String("d", "./", "Directory to serve")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// Get the IP addresses of the current machine
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	// Print the IP addresses
	log.Println("URL:")
	for _, addr := range addrs {
		ip := addr.(*net.IPNet).IP
		if ip.IsLoopback() {
			continue
		}
		log.Printf("http://%s:%d\n", ip, listener.Addr().(*net.TCPAddr).Port)
	}

	err = http.Serve(listener, http.FileServer(http.Dir(*dir)))
	if err != nil {
		log.Fatal(err)
	}
}
