#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·getsp(SB),(NOSPLIT|NOFRAME|WRAPPER),$16-0
	NO_LOCAL_POINTERS
	LEAQ	24(SP), AX
	MOVQ	AX, ret+0(FP)
	RET

TEXT ·getmem(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVQ	addr+0(FP), AX
	MOVQ	0(AX), CX
	MOVQ	CX, ret+8(FP)
	RET
