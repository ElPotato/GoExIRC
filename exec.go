package main

import (
	"os"
	"os/exec"
	"strings"
)

func readCommand(input string) string {
	params := strings.Split(input, " ")
	return params[0]
}

func shellExecute(input string) string {
	cmd, params := splitParams(input, 0, 1, 2)
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

func binaryExecute([]byte) {
	return
}

// POSSIBLE CUT / COPY/PASTE CODE HERE //
func terminate() {
	os.Exit(0)
}