// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·getg_type_v0(SB), NOSPLIT, $0-8
	LEAQ type·main·Goroutine(SB), AX
	MOVQ AX, ret+0(FP)
	RET

TEXT ·getg_type(SB), NOSPLIT, $32-8
	LEAQ type·main·Goroutine(SB), AX
	MOVQ ·gg(SB), BX

	MOVQ AX, 0(SP)
	MOVQ BX, 8(SP)
	CALL runtime·convT2E(SB)
	MOVQ 16(SP), CX

	MOVQ CX, ret+0(FP)
	RET


