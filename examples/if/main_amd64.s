// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·If(SB), NOSPLIT, $0-32
	MOVBQZX ok+8*0(FP), CX // ok
	MOVQ    a+8*1(FP), AX  // a
	MOVQ    b+8*2(FP), BX  // b

	CMPQ    CX, $0         // test ok
	JZ      L              // if !ok, skip 2 line
	MOVQ    AX, ret+24(FP) // return a
	RET

L:
	MOVQ    BX, ret+24(FP) // return b
	RET

