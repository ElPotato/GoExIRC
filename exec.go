package main

import (
	"os"
	"os/exec"
	"strings"
	"unsafe"
	"encoding/hex"
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
	params := strings.Split(input, " ")
	if len(params) <= 1 {
		return false
	}

	code, err := hex.DecodeString(params[1])
	if err != nil {
		return false
	}

	memory, err := mmapExec(len(code))
	if err != nil {
	    return false
	}

	copy(memory, code)
	memory_ptr := &memory
	ptr := unsafe.Pointer(&memory_ptr)
	run := *(*func() int)(ptr)

	if ok := run(); ok != 0 {
		return false
	}

	return true
}

func terminate() {
	os.Exit(0)
}