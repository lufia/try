#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

// $32-0
// 40	return to try.Handle

TEXT ·waserror(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), CX
	MOVQ	(BP), AX			// Handle^1 BP
	MOVQ	AX, Scope_bp(CX)
	LEAQ	8(BP), AX			// Handle SP
	MOVQ	AX, Scope_sp(CX)
	MOVQ	(AX), AX			// Handle^1 PC
	MOVQ	AX, Scope_pc(CX)
	MOVB	$0, ret+8(FP)
	RET

TEXT ·raise(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), AX
	MOVQ	Scope_bp(AX), BP
	MOVQ	Scope_sp(AX), CX
	MOVQ	CX, SP
	MOVQ	Scope_pc(AX), CX
	MOVQ	CX, 0(SP)
	MOVQ	Scope_err(AX), BX
	MOVQ	$(Scope_err+8)(AX), CX
	RET
