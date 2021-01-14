package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/edsrzf/mmap-go"

	"fmt"
	"unsafe"
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
	/*
	package ups

	func run() int {
		return 666
	}

	Binary representation of the code from above
	*/
	code := []byte{
		0x48, 0xc7, 0x44, 0x24, 0x08, 0x00, 0x00, 0x00, 0x00,
		0x48, 0xc7, 0x44, 0x24, 0x08, 0x9a, 0x02, 0x00, 0x00,
		0xc3,
	}

	memory, err := mmap.MapRegion(nil, len(code), mmap.EXEC|mmap.RDWR, mmap.ANON, 0)
	if err != nil {
	    panic(err)
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