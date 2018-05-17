// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h"

TEXT ·main0(SB), $0-0
	RET

TEXT ·main1(SB), $0-8
	RET
TEXT ·main3(SB), $0-24
	RET
TEXT ·main4(SB), $0-32
	RET
TEXT ·main5(SB), $0-40
	RET

// func main()
TEXT ·main2(SB), $16-16
	NO_LOCAL_POINTERS
	MOVQ $1, (SP)
	PCDATA $0, $0
	CALL ·PrintCallerName(SB)
	RET

