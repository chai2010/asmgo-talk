// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// var id int = 9527
DATA  ·id+0(SB)/8,$9527
GLOBL ·id(SB),NOPTR,$8

// func getId() int
TEXT ·getId(SB), $0-8
	MOVQ ·id(SB), AX   // read id
	MOVQ AX, ret+0(FP) // return id
	RET
