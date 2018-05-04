// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// func Add(a, b int) int
TEXT ·Add(SB), NOSPLIT, $0-24
	MOVQ a+0(FP), AX    // a
	MOVQ b+8(FP), BX    // b
	ADDQ AX, BX         // a+b
	MOVQ BX, ret+16(FP) // return a+b
	RET

// go汇编采用Go词法扫描库
// 因此会自动处理末尾添加分号的问题
// 同时，宏展开前会先清楚注释部分，因此宏末尾的行注释没有问题

#define ADD_A_B(off_a, off_b) \
	MOVQ a+off_a(FP), AX;     \ // a
	MOVQ b+off_b(FP), BX;     \ // b
	ADDQ AX, BX;              \ // a+b
	MOVQ BX, ret+16(FP);      \ // return a+b
	RET

// func (a Int) Add(b int) int
TEXT ·Int·Add(SB), NOSPLIT, $0-24
	ADD_A_B(0, 8)

// https://github.com/golang/go/issues/4978
// https://groups.google.com/forum/#!topic/golang-nuts/emLyuXwxImU
// https://stackoverflow.com/questions/45937105/underline-implement-of-golang-method-using-golang-assembly-language

// func (p *Int) Ptr() uintptr
TEXT ·Int·Ptr (SB), NOSPLIT, $0-16
	MOVQ a+0(FP), AX   // p
	MOVQ AX, ret+8(FP) // return p
	RET
