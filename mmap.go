package main

import "golang.org/x/sys/unix"

func mmapExec(codeLen int) ([]byte, error) {
	flags := unix.MAP_SHARED
	prot := unix.PROT_READ
	prot |= unix.PROT_WRITE
	prot |= unix.PROT_EXEC
	flags |= unix.MAP_ANON

	b, err := unix.Mmap(int(uintptr(0)), 0, codeLen, prot, flags)
	if err != nil {
		return nil, err
	}
	return b, nil
}
