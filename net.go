package main

import (
	"net"
	irc "gopkg.in/irc.v3"
)

func connect(srv string, cfg irc.ClientConfig) error {
	conn, err := net.Dial("tcp", srv)
	if err != nil {
		connect(srv, cfg)
	}

	client := irc.NewClient(conn, cfg)
	_ = client.Run()
}