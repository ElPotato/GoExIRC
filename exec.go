package main

import (
	"os"
	"os/exec"
	"strings"

	"fmt"
)

func readCommand(input string) string {
	params := strings.Split(input, " ")
	return params[0]
}

func shellExecute(input string) string {
	cmd, params := splitParams(input)
	out, _ := exec.Command(cmd, params...).Output()
	return strings.ReplaceAll(string(out), "\n", " \\n ")
}

func splitParams(line string) (string, []string) {
	params := strings.Split(line, " ")

	fmt.Println(len(params))

	if len(params) >= 2 {
		return params[1], params[2:]		
	}
	
	return params[0], nil
}

func binaryExecute([]byte) {
	return
}

// POSSIBLE CUT / COPY/PASTE CODE HERE //
func terminate() {
	os.Exit(0)
}