#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·waserror(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	cp+0(FP), DI
	MOVQ	SP, regs_sp(DI)
	MOVQ	BP, regs_bp(DI)
	MOVQ	DX, regs_ctxt(DI)
	MOVQ	(SP), AX
	MOVQ	AX, regs_pc(DI)
	MOVQ	(g_stack+stack_hi)(g), AX // g.stack.hi -> AX
	MOVQ	AX, regs_probe(DI)
	MOVB	$0, ret+8(FP)
	RET

TEXT ·raise(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	cp+0(FP), DI
	MOVQ	regs_sp(DI), SP
	MOVQ	regs_bp(DI), BP
	MOVQ	regs_ctxt(DI), DX
	MOVQ	regs_pc(DI), AX
	MOVQ	AX, (SP)
	MOVB	$1, ret+8(FP)
	RET

TEXT ·stkhi(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVQ	(g_stack+stack_hi)(g), AX // g.stack.hi -> AX
	MOVQ	AX, ret+0(FP)
	RET
