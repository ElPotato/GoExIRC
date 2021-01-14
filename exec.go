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

func shellExecute(cmd string, params []string) string {
	out, _ := exec.Command(cmd, params...).Output()
	return strings.ReplaceAll(string(out), "\n", " \\n ")
}

func splitParams(line string, def, min, max int) (string, []string) {
	params := strings.Split(line, " ")

	if len(params) >= max {
		return params[min], params[max:]		
	}
	
	return params[def], nil
}

func binaryExecute(input string) bool {
	fmt.Println(input)
	return true
}

// POSSIBLE CUT / COPY/PASTE CODE HERE //
func terminate() {
	os.Exit(0)
}