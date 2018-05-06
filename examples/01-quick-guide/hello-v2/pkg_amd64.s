// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

/*
type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type eface struct {
	_type *_type
	data unsafe.Pointer
}

type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}

type interfacetype struct { // 80 bytes on a 64bit arch
	typ     _type
	pkgpath name
	mhdr    []imethod
}

type imethod struct {
	name nameOff
	ityp typeOff
}

type _type struct { // 48 bytes on a 64bit arch
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldalign uint8
	kind       uint8
	alg        *typeAlg
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}

func resolveNameOff(ptrInModule unsafe.Pointer, off nameOff) name {}
func resolveTypeOff(ptrInModule unsafe.Pointer, off typeOff) *_type {}
*/

TEXT ·HelloWorld(SB),$56-0
	// var a0 interface{} = helloworld
	// MOVQ $0, a0_type-8*2(SP) // a0._type
	// MOVQ $0, a0_data-8*1(SP) // a0.data

	// func runtime.convT2E(t *_type, elem unsafe.Pointer) (e eface)
	LEAQ type·string(SB), AX; MOVQ AX, 8*0(SP) // arg:t: string type
	LEAQ ·helloworld(SB), BX; MOVQ BX, 8*1(SP) // arg:elem: &helloworld
	CALL runtime·convT2E(SB)
	MOVQ 8*2(SP), AX; MOVQ AX, a0_type-8*2(SP) // ret:e._type
	MOVQ 8*3(SP), BX; MOVQ BX, a0_data-8*1(SP) // ret:e.data

	// func fmt.Println(a ...interface{}) (n int, err error)
	LEAQ a0_type-8*2(SP), AX; MOVQ AX, 8*0(SP) // arg:a.Data
	MOVQ $1, 8*1(SP)                           // arg:a.Len
	MOVQ $1, 8*2(SP)                           // arg:a.Cap
	CALL fmt·Println(SB)                       // ret: ignore

	// MOVQ 8*3(SP), AX                        // ret:n
	// MOVQ 8*4(SP), BX                        // ret:err._type
	// MOVQ 8*5(SP), CX                        // ret:err.data

	RET
