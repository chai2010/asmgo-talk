// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.9

#include "textflag.h"
#include "funcdata.h"

#define G_GOID_OFFSET 152 // go1.9,go1.10

// func get_g_id() int64
TEXT ·get_g_id(SB), NOSPLIT, $0-8

	// get runtime.g
	MOVQ TLS, CX
	MOVQ 0(CX)(TLS*1), AX

	// read AX[G_GOID_OFFSET]
	MOVQ $G_GOID_OFFSET, BX
	MOVQ 0(AX)(BX*1), CX

	// return goid
	MOVQ CX, ret+0(FP)
	RET

	LEAQ 0(AX)(BX*1), DX
	MOVQ (DX), AX
	MOVQ AX, ret+0(FP)
	RET
