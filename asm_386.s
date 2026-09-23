#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·waserror(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVL	cp+0(FP), DI
	MOVL	SP, regs_sp(DI)
	MOVL	DX, regs_ctxt(DI)
	MOVL	(SP), AX
	MOVL	AX, regs_pc(DI)
	MOVL	(TLS), CX
	MOVL	(g_stack+stack_hi)(CX), AX
	MOVL	AX, regs_probe(DI)
	MOVB	$0, ret+4(FP)
	RET

TEXT ·raise(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVL	cp+0(FP), DI
	MOVL	regs_sp(DI), SP
	MOVL	regs_ctxt(DI), DX
	MOVL	regs_pc(DI), AX
	MOVL	AX, (SP)
	MOVB	$1, ret+4(FP)
	RET

TEXT ·stkhi(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVL	(TLS), DI
	MOVL	(g_stack+stack_hi)(DI), AX // g.stack.hi -> AX
	MOVL	AX, ret+0(FP)
	RET
