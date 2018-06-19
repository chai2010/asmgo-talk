package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func getg_type_v0() unsafe.Pointer
func getg_type() unsafe.Pointer

// https://golang.org/src/reflect/type.go#L1411
// https://golang.org/src/reflect/value.go?h=emptyInterface#L181
// https://golang.org/src/reflect/type.go?h=toType#L3040
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

var gg Goroutine

func main() {
	t0 := typeOfG()
	tye0 := (*emptyInterface)(unsafe.Pointer(&t0))
	fmt.Printf("tye0: %#v\n", tye0)

	v1 := Goroutine{}
	fmt.Printf("v1: %p\n", &v1)

	t1 := reflect.TypeOf(v1)
	tye1 := (*emptyInterface)(unsafe.Pointer(&t1))
	fmt.Printf("tye1: %#v\n", tye1)
}

//go:noinline
//go:nosplit
func Debug() interface{} {
	return Goroutine{}
}

func typeOfG() reflect.Type {
	return typeOf(getg_type())
}

func typeOf(_type unsafe.Pointer) reflect.Type {

	var x reflect.Type
	var px = (*emptyInterface)(unsafe.Pointer(&x))

	fmt.Println("_type:", _type)
	px.typ = _type
	//px.word = getg_addr()
	return x
}

type stack struct {
	lo uintptr
	hi uintptr
}

type Goroutine struct {
	// Stack parameters.
	// stack describes the actual stack memory: [stack.lo, stack.hi).
	// stackguard0 is the stack pointer compared in the Go stack growth prologue.
	// It is stack.lo+StackGuard normally, but can be StackPreempt to trigger a preemption.
	// stackguard1 is the stack pointer compared in the C stack growth prologue.
	// It is stack.lo+StackGuard on g0 and gsignal stacks.
	// It is ~0 on other goroutine stacks, to trigger a call to morestackc (and crash).
	stack       stack   // offset known to runtime/cgo
	stackguard0 uintptr // offset known to liblink
	stackguard1 *int    // offset known to liblink
}
