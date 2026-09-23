#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·waserror(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVD	cp+0(FP), R2
	MOVD	RSP, R0
	MOVD	R0, regs_sp(R2)
	MOVD	R29, regs_bp(R2)
	MOVD	R26, regs_ctxt(R2)
	MOVD	LR, regs_pc(R2)
	MOVD	(g_stack+stack_hi)(g), R0 // g.stack.hi -> R0
	MOVD	R0, regs_probe(R2)
	MOVD	$0, R0
	MOVD	R0, ret+8(FP)
	RET

TEXT ·raise(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVD	cp+0(FP), R2
	MOVD	regs_sp(R2), R0
	MOVD	R0, RSP
	MOVD	regs_bp(R2), R29
	MOVD	regs_ctxt(R2), R26
	MOVD	regs_pc(R2), LR
	MOVD	$1, R0
	MOVD	R0, ret+8(FP)
	RET

TEXT ·stkhi(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVD	(g_stack+stack_hi)(g), R0 // g.stack.hi -> R0
	MOVD	R0, ret+0(FP)
	RET
