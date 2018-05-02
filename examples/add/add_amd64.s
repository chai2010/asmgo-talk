// Copyright © 2017 ChaiShushan <chaishushan{AT}gmail.com>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

#include "textflag.h"

// func Add(a, b int) int
TEXT ·Add(SB), NOSPLIT, $0-24
	MOVQ a+0(FP), AX    // a
	MOVQ b+8(FP), BX    // b
	ADDQ AX, BX         // a+b
	MOVQ BX, ret+16(FP) // return a+b
	RET

#define ADD_A_B(off_a, off_b) \
	MOVQ a+off_a(FP), AX      \ // a
	MOVQ b+off_b(FP), BX      \ // b
	ADDQ AX, BX               \ // a+b
	MOVQ BX, ret+16(FP)       \ // return a+b
	RET

// func (a Int) Add(b int) int
TEXT ·Int·Add(SB), NOSPLIT, $0-24
	ADD_A_B(0, 8)

// https://groups.google.com/forum/#!topic/golang-nuts/emLyuXwxImU

// func (p *Int) Ptr() uintptr
TEXT ·Int·Ptr (SB), NOSPLIT, $0-16
	MOVQ a+0(FP), AX   // p
	MOVQ AX, ret+8(FP) // return p
	RET
