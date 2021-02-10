package main

import (
	"net"
	irc "gopkg.in/irc.v3"
)

func connect(srv string, cfg irc.ClientConfig) {
	conn, err := net.Dial("tcp", srv)
	if err != nil {
		connect(srv, cfg)
	}

	if err := client(conn, cfg); err != nil {
		// nolint
		client(conn, cfg)
	}
}

func client(conn net.Conn, cfg irc.ClientConfig) error {
	client := irc.NewClient(conn, cfg)
	if err := client.Run(); err != nil {
		return err
	}
	
	return nil
}