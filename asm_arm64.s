#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·waserror(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVD	s+0(FP), R2
	MOVD	RSP, R1
	MOVD	R1, Scope_sp(R2)
	MOVD	R30, Scope_pc(R2)
	MOVD	$0, R0
	MOVD	R0, ret+8(FP)
	RET

TEXT ·raise(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVD	s+0(FP), R2
	MOVD	Scope_sp(R2), R1
	MOVD	R1, RSP
	MOVD	Scope_pc(R2), R30
	MOVD	$1, R0
	MOVD	R0, ret+8(FP)
	RET
