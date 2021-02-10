package main

import irc "gopkg.in/irc.v3"

const (
	SRVADDR string = "chat.freenode.net:6667"
	CHANNEL string = "#yyyzzzxxx"
	PASS    string = "password"
	UNAME   string = "username"
	FNAME   string = "Full Name"
)

var ID string

func main() {
	connect(SRVADDR, irc.ClientConfig{Nick: ID, Pass: PASS, User: UNAME, Name: FNAME,
		Handler: irc.HandlerFunc(handler)})
}
