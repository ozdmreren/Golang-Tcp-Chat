package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type client struct {
	name string
	conn net.Conn
	msg  string
}

func (c *client) ReadInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')

		if err != nil {
			fmt.Println(err.Error())
		}

		c.msg = strings.TrimSpace(msg)

		c.conn.Write([]byte(fmt.Sprintf("%s:  | Enter your message: ", c.conn.RemoteAddr().String())))

		if c.msg == "/exit" {
			fmt.Printf("%s left the server. Time: %s", c.name, time.Now().String())
			c.conn.Close()

			return
		}

		fmt.Printf("%s: %s\n", c.name, c.msg)
	}
}
