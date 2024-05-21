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

// func startServer() {
// 	fs := http.FileServer(http.Dir("./"))
// 	http.Handle("/", fs)

// 	//log.Print("Listening on :8888...")
// 	err := http.ListenAndServe(":0", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
func startServer() {
	port := flag.Int("p", 0, "Port number to use, default is 0 for random")
	dir := flag.String("d", "./", "Directory to serve")
	flag.Parse()

	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/", fs)

	listener, err := net.Listen("tcp", fmt.Sprint(":", *port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("Listening on %s...", listener.Addr())
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
}
