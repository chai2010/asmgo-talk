// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h"

// func sum(n int) int
TEXT ·sum(SB), $32-16
	MOVQ n+0(FP), AX       // n
	MOVQ result+8(FP), BX  // result

	MOVQ $0, temp2-8*1(SP) // DX: temp2
	MOVQ $0, temp1-8*2(SP) // CX: temp1

	CMPQ AX, $0            // test n - 0
	JLE  L_END             // if <= 0: goto goto LEND
	// JMP  L_STEP_TO_END  // goto L_STEP_TO_END

L_STEP_TO_END:
	MOVQ AX, CX            // CX: temp1 = n
	ADDQ $-1, CX           // CX: temp1 += -1
	MOVQ CX, temp1-8*2(SP)

	MOVQ CX, 0(SP)         // arg: n-1
	CALL ·sum(SB)
	MOVQ 8(SP), BX         // DX: temp2 = sum(n-1)

	MOVQ n+0(FP), AX       // n
	ADDQ AX, BX            // DX: temp2 += n
	MOVQ BX, result+8(FP)  // return result
	RET

L_END:
	MOVQ $0, result+8(FP) // return 0
	RET


// func sum(n int) int
TEXT ·sum_v2(SB), $32-16
	NO_LOCAL_POINTERS

L_START:
	MOVQ TLS, CX
	MOVQ 0(CX)(TLS*1), AX
	CMPQ SP, 16(AX)
	JLS  L_MORE_STK

	MOVQ n+0(FP), AX       // n
	MOVQ result+8(FP), BX  // result

	MOVQ $0, temp2-8*1(SP) // DX: temp2
	MOVQ $0, temp1-8*2(SP) // CX: temp1

	CMPQ AX, $0            // test n - 0
	JLE  L_END             // if <= 0: goto goto LEND
	// JMP  L_STEP_TO_END  // goto L_STEP_TO_END

L_STEP_TO_END:
	MOVQ AX, CX            // CX: temp1 = n
	ADDQ $-1, CX           // CX: temp1 += -1
	MOVQ CX, temp1-8*2(SP)

	MOVQ CX, 0(SP)         // arg: n-1
	CALL ·sum_v2(SB)
	MOVQ 8(SP), BX         // DX: temp2 = sum(n-1)

	MOVQ n+0(FP), AX       // n
	ADDQ AX, BX            // DX: temp2 += n
	MOVQ BX, result+8(FP)  // return result
	RET

L_END:
	MOVQ $0, result+8(FP) // return 0
	RET

L_MORE_STK:
	CALL runtime·morestack_noctxt(SB)
	JMP  L_START
