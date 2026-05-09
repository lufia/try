#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·waserror(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVL	cp+0(FP), DI
	MOVL	SP, Checkpoint_sp(DI)
	MOVL	DX, Checkpoint_ctxt(DI)
	MOVL	(SP), AX
	MOVL	AX, Checkpoint_pc(DI)
	MOVL	(TLS), CX
	MOVL	(g_stack+stack_hi)(CX), AX
	MOVL	AX, Checkpoint_probe(DI)
	MOVB	$0, ret+4(FP)
	RET

TEXT ·raise(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVL	cp+0(FP), DI
	MOVL	Checkpoint_sp(DI), SP
	MOVL	Checkpoint_ctxt(DI), DX
	MOVL	Checkpoint_pc(DI), AX
	MOVL	AX, (SP)
	MOVB	$1, ret+4(FP)
	RET

// returns stack-top
TEXT ·getbp(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVL	(TLS), DI
	MOVL	(g_stack+stack_hi)(DI), AX // g.stack.hi -> AX
	MOVL	AX, ret+4(FP)
	RET
