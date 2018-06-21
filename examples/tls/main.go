package main

import "unsafe"

func get_tls() (tls0, tls1 unsafe.Pointer)

func main() {
	tls0, tls1 := get_tls()
	println(tls0, tls1)
}
