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
	RET

TEXT ·raise(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), CX
	MOVQ	Scope_sp(CX), SP
	MOVQ	Scope_bp(CX), BP
	MOVQ	Scope_pc(CX), AX
	MOVQ	AX, 0(SP)
	RET

TEXT ·raise1(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), CX
	MOVQ	skip+8(FP), AX
	MOVQ	BP, BX		// Raise BP
loop:
	CMPQ	AX, $0
	JZ		next
	MOVQ	(BX), BX	// Raise^skip BP
	DECQ	AX
	JMP		loop

next:
	MOVQ	Scope_pc(CX), AX
	MOVQ	AX, 8(BX)
	MOVQ	CX, AX		// first arg
	MOVQ	Scope_err(CX), BX
	MOVQ	$(Scope_err+8)(CX), CX
	RET
