// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://github.com/golang/go/blob/master/src/runtime/runtime2.go
// type g struct { ... }

package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/cch123/goroutineid"
	"github.com/petermattis/goid"
)

func getg() interface{}

func getg_type() unsafe.Pointer
func getg_addr() unsafe.Pointer

func main() {
	//go func() {
	//	fmt.Println("-------------")
	//	fmt.Printf("other goid: %0X\n", GetGoID())
	//	fmt.Println("-------------")
	//
	//	fmt.Printf("main getg: %v\n", getg())
	//}()

	go func() {
		fmt.Println("-------------")
		gid0 := goroutineid.GetGoID()
		fmt.Printf("main gid0: %0X(%d)\n", gid0, gid0)

		gid1 := goid.Get()
		fmt.Printf("main gid1: %0X(%d)\n", gid1, gid1)

		fmt.Println("-------------")
		//fmt.Printf("main getg_type: %0X\n", getg_type())
		fmt.Printf("main getg_addr000: %0X\n", getg_addr())

		p := (*g)(getg_addr())
		fmt.Printf("main getg_addr: %p(%d)\n", p, uintptr(unsafe.Pointer(p)))

		fmt.Printf("getg_addr goid offsetof: %d, %d\n", unsafe.Offsetof(p.goid), offsetDict["go1.10"])

		fmt.Printf("getg_addr v.id: %0X(%d)\n", p.goid, p.goid)
	}()
	//fmt.Printf("main getg: %v\n", getg())

	time.Sleep(time.Second)
}

type stack struct {
	lo uintptr
	hi uintptr
}

type g struct {
	// Stack parameters.
	// stack describes the actual stack memory: [stack.lo, stack.hi).
	// stackguard0 is the stack pointer compared in the Go stack growth prologue.
	// It is stack.lo+StackGuard normally, but can be StackPreempt to trigger a preemption.
	// stackguard1 is the stack pointer compared in the C stack growth prologue.
	// It is stack.lo+StackGuard on g0 and gsignal stacks.
	// It is ~0 on other goroutine stacks, to trigger a call to morestackc (and crash).
	stack       stack   // offset known to runtime/cgo
	stackguard0 uintptr // offset known to liblink
	stackguard1 uintptr // offset known to liblink

	_panic       unsafe.Pointer // innermost panic - offset known to liblink
	_defer       unsafe.Pointer // innermost defer
	m            unsafe.Pointer // current m; offset known to arm liblink
	sched        gobuf
	syscallsp    uintptr        // if status==Gsyscall, syscallsp = sched.sp to use during gc
	syscallpc    uintptr        // if status==Gsyscall, syscallpc = sched.pc to use during gc
	stktopsp     uintptr        // expected sp at top of stack, to check in traceback
	param        unsafe.Pointer // passed parameter on wakeup
	atomicstatus uint32
	stackLock    uint32 // sigprof/scang lock; TODO: fold in to atomicstatus
	goid         int64
}

type gobuf struct {
	// The offsets of sp, pc, and g are known to (hard-coded in) libmach.
	//
	// ctxt is unusual with respect to GC: it may be a
	// heap-allocated funcval, so GC needs to track it, but it
	// needs to be set and cleared from assembly, where it's
	// difficult to have write barriers. However, ctxt is really a
	// saved, live register, and we only ever exchange it between
	// the real register and the gobuf. Hence, we treat it as a
	// root during stack scanning, which means assembly that saves
	// and restores it doesn't need write barriers. It's still
	// typed as a pointer so that any other writes from Go get
	// write barriers.
	sp   uintptr
	pc   uintptr
	g    guintptr
	ctxt unsafe.Pointer
	ret  sys_Uintreg
	lr   uintptr
	bp   uintptr // for GOEXPERIMENT=framepointer
}

type guintptr uintptr
type sys_Uintreg uint64
