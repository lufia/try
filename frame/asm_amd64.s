#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·getsp(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVQ	SP, ret+0(FP)
	RET

TEXT ·getfp(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVQ	BP, ret+0(FP)
	RET

TEXT ·returnTo(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVQ	off+0(FP), AX
	MOVQ	SP, BX
	ADDQ	AX, BX
	MOVQ	0(BX), CX
	MOVQ	CX, ret+8(FP)
	RET

// 24 = 8(return) + 8(off) + 8(ret)
#define bp_off 40

TEXT ·returnTo2(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVQ	off+0(FP), AX
	MOVQ	bp_off(SP), BX
	ADDQ	AX, BX
	MOVQ	0(BX), CX
	MOVQ	CX, ret+8(FP)
	RET
