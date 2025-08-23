#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·waserror(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), CX
	MOVQ	SP, Scope_sp(CX)
	MOVQ	0(SP), AX
	MOVQ	AX, Scope_pc(CX)
	MOVQ	$0, AX
	RET

TEXT ·raise(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), CX
	MOVQ	Scope_sp(CX), SP
	MOVQ	Scope_pc(CX), AX
	MOVQ	AX, 0(SP)
	MOVQ	$1, AX
	RET
