package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func getg_type_v0() unsafe.Pointer
func getg_type(data *Goroutine) interface{}

// https://golang.org/src/reflect/type.go#L1411
// https://golang.org/src/reflect/value.go?h=emptyInterface#L181
// https://golang.org/src/reflect/type.go?h=toType#L3040
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

var AA int
var gg Goroutine

var gg111 bool
var gh bool
var x bool
var g bool

var pgg = &gg

func get_gg_ptr() *Goroutine

func main() {
	gg.stackguard0 = 9527

	fmt.Printf("gg0: %p\n", &gg)
	fmt.Printf("gg0: %p\n", get_gg_ptr())

	t0 := typeOfG()
	tye0 := (*emptyInterface)(unsafe.Pointer(&t0))
	fmt.Printf("tye0: %#v\n", tye0)

	if field, ok := t0.FieldByName("stackguard0"); ok {
		fmt.Println("Name:", field.Name)
		fmt.Println("Offset:", field.Offset)
	}

	x := getg_type(&gg).(Goroutine).stackguard0
	println(x)
	return

	v1 := Goroutine{}
	fmt.Printf("v1: %p\n", &v1)

	t1 := reflect.TypeOf(v1)
	tye1 := (*emptyInterface)(unsafe.Pointer(&t1))
	fmt.Printf("tye1: %#v\n", tye1)
}

//go:noinline
//go:nosplit
func CallDebug() {
	var x = Debug()
	_ = x
}

//go:noinline
//go:nosplit
func Debug() interface{} {
	return Goroutine{}
}

func typeOfG() reflect.Type {
	return reflect.TypeOf(getg_type(&gg))
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
