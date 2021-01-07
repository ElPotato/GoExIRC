package main

import (
	"os/exec"
	"strings"
)

func executeCommand(input string) string {
	cmd, params := splitParams(input)
	out, _ := exec.Command(cmd, params...).Output()
	return strings.ReplaceAll(string(out), "\n", " \\n ")
}

func splitParams(line string) (string, []string) {
	params := strings.Split(line, " ")
	return params[0], params[1:]
}