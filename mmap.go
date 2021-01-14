package main

import "golang.org/x/sys/unix"

const (
	ANON = 1 << iota; RDWR = ANON
	EXEC
)

func mmapExec(codeLen int, inprot, inflags uintptr, off int64) ([]byte, error) {
	flags := unix.MAP_SHARED
	prot := unix.PROT_READ

	if inprot&RDWR != 0 {
		prot |= unix.PROT_WRITE
	}

	if inprot&EXEC != 0 {
		prot |= unix.PROT_EXEC
	}

	if inflags&ANON != 0 {
		flags |= unix.MAP_ANON
	}

	b, err := unix.Mmap(int(uintptr(0)), off, codeLen, prot, flags)
	if err != nil {
		return nil, err
	}
	return b, nil
}