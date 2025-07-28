package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	s := CreateNewServer()
	listener, err := net.Listen("tcp", ":8888")

	fmt.Printf("Server started at %s\n", "8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	// Accept Loop
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go s.CreateNewClient(conn)
	}

}
