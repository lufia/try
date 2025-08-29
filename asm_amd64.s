#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

// $32-0
#define	Handle_size 32

TEXT ·waserror(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), CX
	MOVQ	32(SP), AX
	MOVQ	(AX), AX
	ADDQ	$8, AX
	MOVQ	AX, Scope_sp(CX)
	MOVQ	40(SP), AX
	MOVQ	AX, Scope_pc(CX)
	MOVB	$0, ret+8(FP)
	RET

TEXT ·raise(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), AX
	MOVQ	Scope_sp(AX), BP
	SUBQ	$8, BP
	MOVQ	Scope_sp(AX), SP
	MOVQ	Scope_pc(AX), CX
	MOVQ	CX, 0(SP)
	MOVQ	Scope_err(AX), BX
	MOVQ	$(Scope_err+8)(AX), CX
	RET
