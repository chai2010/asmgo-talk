// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h"
#include "go_tls.h"

// func getg_type() unsafe.Pointer
TEXT ·getg_type(SB), NOSPLIT, $0-8
	MOVQ $type·runtime·g(SB), AX
	MOVQ AX, ret+0(FP)
	RET

// func getg_addr() unsafe.Pointer
TEXT ·getg_addr(SB), NOSPLIT, $0-8
	get_tls(CX)
	MOVQ g(CX), AX
	MOVQ AX, ret+0(FP)
	RET

// func getg() interface{}
TEXT ·getg(SB), NOSPLIT, $32-16
	MOVQ $0, ret+0(FP)
	MOVQ $0, ret+8(FP)
	GO_RESULTS_INITIALIZED

	// get runtime.g type
	MOVQ $type·runtime·g(SB), AX

	// get runtime.g
	get_tls(CX);
	MOVQ g(CX), BX;

	// convert (*g) to interface{}
	MOVQ AX, 0(SP)
	MOVQ BX, 8(SP)
	CALL runtime·convT2E(SB)
	MOVQ 16(SP), AX
	MOVQ 24(SP), BX

	// return interface{}
	MOVQ AX, ret+0(FP)
	MOVQ BX, ret+8(FP)
	RET
