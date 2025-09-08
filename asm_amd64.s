#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

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
	MOVQ	s+0(FP), CX
	MOVQ	Scope_bp(CX), AX
	MOVQ	(BP), BX	// Raise^1 BP
loop:
	CMPQ	AX, BX
	JEQ		next
	MOVQ	(BX), BX
	JMP		loop
next:
	MOVQ	Scope_pc(CX), AX
	MOVQ	AX, 8(BX)
	RET

TEXT ·rewind(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), CX
	MOVQ	Scope_err(CX), BX
	MOVQ	$(Scope_err+8)(CX), CX
	RET
