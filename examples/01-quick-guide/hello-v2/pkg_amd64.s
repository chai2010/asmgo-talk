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

type _type struct {
	name string
	bits uint
	signed bool
}

type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}
*/

TEXT ·HelloWorld(SB),$48-0
	// func runtime.convT2E(t *_type, elem unsafe.Pointer) (e eface)
	LEAQ type·string(SB),   AX; MOVQ AX, 0(SP) // arg:t: string type
	LEAQ ·helloworld+0(SB), AX; MOVQ AX, 8(SP) // arg:elem: &helloworld
	CALL runtime·convT2E(SB)
	MOVQ 16(SP), AX;                           // ret:e {_type, data}

	// func fmt.Println(a ...interface{}) (n int, err error)
	MOVQ AX, 0(SP)       // arg:a.Data
	MOVQ $1, 8(SP)       // arg:a.Len
	MOVQ $1, 16(SP)      // arg:a.Cap
	CALL fmt·Println(SB) // ret: ignore

	RET
