
#include "textflag.h"
#include "go_tls.h"

// func get_tls() (tls0, tls1 unsafe.Pointer)
TEXT ·get_tls(SB), NOSPLIT, $0-16
	MOVQ TLS, AX
	get_tls(BX)
	MOVQ AX, ret0+0(FP)
	MOVQ BX, ret1+8(FP)
	RET
