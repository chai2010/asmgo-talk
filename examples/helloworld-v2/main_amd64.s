// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h"

// func main()
TEXT ·main(SB), $16-0
	MOVQ ·helloworld(SB), AX
	MOVQ AX,  str_data-16(SP)
	MOVQ $16, str_len-8(SP)
	CALL runtime·printstring(SB)
	RET
