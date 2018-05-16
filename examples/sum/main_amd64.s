// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// func sum(n int) int
TEXT ·sum(SB), $16-16
	MOVQ n+0(FP), AX       // n
	MOVQ result+8(FP), BX  // result

	MOVQ $0, temp2-8*1(SP) // DX: temp2
	MOVQ $0, temp1-8*2(SP) // CX: temp1

	CMPQ AX, $0           // test n - 0
	JLE  L_END            // if <= 0: goto goto LEND
	JMP  L_STEP_TO_END    // goto L_STEP_TO_END

L_STEP_TO_END:
	MOVQ AX, CX           // CX: temp1 = n
	ADDQ $-1, CX          // CX: temp1 += -1

	MOVQ CX, 0(SP)        // arg: n-1
	CALL ·sum(SB)
	MOVQ 8(SP), DX        // DX: temp2 = sum(n-1)

	ADDQ AX, DX           // DX: temp2 += n
	MOVQ DX, result+8(FP) // return result
	RET

L_END:
	MOVQ $0, result+8(FP) // return 0
	RET
