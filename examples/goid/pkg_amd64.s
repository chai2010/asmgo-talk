// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "go_tls.h"

// func GetGoID() int64
TEXT ·GetGoID(SB), NOSPLIT, $0-8
	get_tls(CX)
	MOVQ g(CX), AX
	MOVQ ·offset(SB), BX
	LEAQ 0(AX)(BX*1), DX
	MOVQ (DX), AX
	MOVQ AX, ret+0(FP)
	RET

TEXT ·getg_type(SB), NOSPLIT, $0-8
	LEAQ type·runtime·g(SB), AX
	MOVQ AX, ret+0(FP)
	RET

TEXT ·getg_addr(SB), NOSPLIT, $0-8
	get_tls(CX)
	MOVQ g(CX), AX
	MOVQ AX, ret+0(FP)
	RET

TEXT ·getg(SB),$32-16

L_START:
	MOVQ TLS, CX
	MOVQ 0(CX)(TLS*1), AX
	CMPQ SP, 16(AX)
	JLS  L_MORE_STK

	// get runtime.g type
	LEAQ type·runtime·g(SB), AX;

	// get runtime.g
	get_tls(CX);
	MOVQ g(CX), BX


	//LEAQ go·itab·runtime·g(SB), BX;
	//MOVQ BX, ret_type-8*0(FP)
	//MOVQ AX, ret_type-8*1(FP)


	MOVQ AX, 0(SP)
	MOVQ BX, 8(SP)
	CALL runtime·convT2E(SB)
	MOVQ 16(SP), AX
	MOVQ 24(SP), BX

	MOVQ AX, ret+0(FP)
	MOVQ BX, ret+8(FP)

	RET

L_MORE_STK:
	CALL runtime·morestack_noctxt(SB)
	JMP  L_START
