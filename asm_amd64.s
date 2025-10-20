#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·waserror(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	cp+0(FP), DI
	MOVQ	SP, Checkpoint_sp(DI)
	MOVQ	BP, Checkpoint_bp(DI)
	MOVQ	DX, Checkpoint_ctxt(DI)
	MOVQ	(SP), AX
	MOVQ	AX, Checkpoint_pc(DI)
	MOVQ	(BP), AX
	MOVQ	AX, Checkpoint_probe(DI)
	MOVB	$0, ret+8(FP)
	RET

TEXT ·raise(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	cp+0(FP), DI
	MOVQ	Checkpoint_sp(DI), SP
	MOVQ	Checkpoint_bp(DI), BP
	MOVQ	Checkpoint_ctxt(DI), DX
	MOVQ	Checkpoint_pc(DI), AX
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
