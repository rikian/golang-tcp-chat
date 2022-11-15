package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()

	listener, err := net.Listen("tcp", ":9091")

	if err != nil {
		log.Fatalf("unable to start server, error: %v", err.Error())
	}

	defer listener.Close()

	log.Print("server listen on port 9091")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Panicf("unable to accept connection, error: %v", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
