#include "textflag.h"
#include "funcdata.h"

// func main()
TEXT ·main(SB), $16-0
	NO_LOCAL_POINTERS
	MOVQ ·gopkgHelloWorld(SB), AX
	MOVQ AX, (SP)
	MOVQ $16, 8(SP)
	CALL runtime·printstring(SB)
	RET
