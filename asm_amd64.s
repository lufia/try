#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·waserror(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), CX
	MOVQ	SP, Scope_sp(CX)
	MOVQ	(SP), AX
	MOVQ	AX, Scope_pc(CX)
	MOVB	$0, ret+8(FP)
	RET

TEXT ·raise(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), CX
	MOVQ	Scope_sp(CX), SP
	MOVQ	Scope_pc(CX), AX
	MOVQ	AX, (SP)
	MOVB	$1, ret+8(FP)
	RET
