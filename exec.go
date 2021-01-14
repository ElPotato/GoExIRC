package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/edsrzf/mmap-go"

	"fmt"
	"unsafe"
	// "regexp"
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
	code, _ := hex.DecodeString(params[1])

	memory, err := mmap.MapRegion(nil, len(code), mmap.EXEC|mmap.RDWR, mmap.ANON, 0)
	if err != nil {
	    return false
	}

	copy(memory, code)

	memory_ptr := &memory
	ptr := unsafe.Pointer(&memory_ptr)
	run := *(*func() int)(ptr)

	fmt.Println(run())

	return true
}

// POSSIBLE CUT / COPY/PASTE CODE HERE //
func terminate() {
	os.Exit(0)
}