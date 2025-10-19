#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·waserror(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), DI
	MOVQ	SP, Scope_sp(DI)
	MOVQ	BP, Scope_bp(DI)
	MOVQ	DX, Scope_dx(DI)
	MOVQ	(SP), AX
	MOVQ	AX, Scope_pc(DI)
	MOVQ	(BP), AX
	MOVQ	AX, Scope_probe(DI)
	MOVB	$0, ret+8(FP)
	RET

TEXT ·raise(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), DI
	MOVQ	Scope_sp(DI), SP
	MOVQ	Scope_bp(DI), BP
	MOVQ	Scope_dx(DI), DX
	MOVQ	Scope_pc(DI), AX
	MOVQ	AX, (SP)
	MOVB	$1, ret+8(FP)
	RET

TEXT ·getbp(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	skip+0(FP), DI
	MOVQ	BP, AX
loop:
	CMPQ	DI, $0
	JBE	end
	MOVQ	(AX), AX
	DECQ	DI
	JMP	loop
end:
	MOVQ	AX, ret+8(FP)
	RET
