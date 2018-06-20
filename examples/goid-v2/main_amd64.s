// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// func getg_addr() unsafe.Pointer
TEXT ·getg_addr(SB), NOSPLIT, $0-8
	MOVQ TLS, CX
	MOVQ 0(CX)(TLS*1), AX
	MOVQ AX, ret+0(FP)
	RET
