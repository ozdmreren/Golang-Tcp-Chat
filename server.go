package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type server struct {
	clients []client
}

func (s *server) CreateNewClient(conn net.Conn) {
	c := &client{conn: conn}
	s.clients = append(s.clients, *c)

	c.conn.Write([]byte("Enter your chat nick: "))
	reader := bufio.NewReader(c.conn)
	c.name, _ = reader.ReadString('\n')
	c.name = strings.TrimSpace(c.name)

	fmt.Printf("%s join the server.\n", c.name)

	go c.ReadInput()
}

func CreateNewServer() *server {
	return &server{
		clients: make([]client, 10),
	}
}
