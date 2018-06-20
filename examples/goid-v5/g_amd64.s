// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h"

// func GetGroutinePointer() uintptr
TEXT ·GetGroutinePointer(SB), NOSPLIT, $0-8
	MOVQ TLS, CX
	MOVQ 0(CX)(TLS*1), AX
	MOVQ AX, ret+0(FP)
	RET

// func GetGroutine() interface{}
TEXT ·GetGroutine(SB), NOSPLIT, $32-16
	MOVQ $0, ret+0(FP)
	MOVQ $0, ret+8(FP)
	GO_RESULTS_INITIALIZED

	// get runtime.g
	MOVQ TLS, CX
	MOVQ 0(CX)(TLS*1), AX

	// get runtime.g type
	MOVQ $type·runtime·g(SB), BX

	// convert (*g) to interface{}
	MOVQ AX, 8(SP)
	MOVQ BX, 0(SP)
	CALL runtime·convT2E(SB)
	MOVQ 16(SP), AX
	MOVQ 24(SP), BX

	// return interface{}
	MOVQ AX, ret+0(FP)
	MOVQ BX, ret+8(FP)
	RET
