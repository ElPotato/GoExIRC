# GoExIRC
A small node which execute bash commands received VIA IRC.


## To Do

* resolve auto-start/ systemd problem
* add files transport
* encryption?
* resolve private message communication
* add auth to "terminate" command // add remove binary/ trails
* pre-processor/ binary obfuscation

## How to use

`sh ls -lha`
`bin 48c74424080000000048c744240800000000c3`

To get hex of compiled go code:

```
package works

func run() int {
	return 0
}
```

`go tool compile -S -N file.go`
`go tool objdump -S file.go`

```
func run() int {
  0x264			48c744240800000000	MOVQ $0x0, 0x8(SP)
	return 1
  0x26d			48c744240801000000	MOVQ $0x1, 0x8(SP)
  0x276			c3			RET
  ```

Take only hex data e.g. `48c744240800000000`, `48c744240801000000` , `c3`

Function in byte code must return 0 for OK status and 1 for ERR.

## Tools
[Shellgo - Extract shellcode from objdump](https://github.com/ElPotato/shellgo)