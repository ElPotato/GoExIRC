package main

import (
	"net"
	irc "gopkg.in/irc.v3"
)

func connect(srv string, cfg irc.ClientConfig) {
	conn, _ := net.Dial("tcp", srv) // retry here

	client := irc.NewClient(conn, cfg)
	_ = client.Run()
}