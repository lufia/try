#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

TEXT ·waserror(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVD	s+0(FP), R2
	MOVD	RSP, R0
	MOVD	R0, Scope_sp(R2)
	MOVD	R29, Scope_bp(R2)
	MOVD	R26, Scope_ctxt(R2)
	MOVD	LR, Scope_pc(R2)
	MOVD	(R29), R0
	MOVD	R0, Scope_probe(R2)
	MOVD	$0, R0
	MOVD	R0, ret+8(FP)
	RET

TEXT ·raise(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVD	s+0(FP), R2
	MOVD	Scope_sp(R2), R0
	MOVD	R0, RSP
	MOVD	Scope_bp(R2), R29
	MOVD	Scope_ctxt(R2), R26
	MOVD	Scope_pc(R2), LR
	MOVD	$1, R0
	MOVD	R0, ret+8(FP)
	RET

TEXT ·getbp(SB),NOSPLIT|NOFRAME,$0
	NO_LOCAL_POINTERS
	MOVD	skip+0(FP), R2
	MOVD	R29, R0
loop:
	CMP	$0, R2 // 0 > R2
	BLE	end
	MOVD	(R0), R0
	SUB	$1, R2, R2
	JMP	loop
end:
	MOVD	R0, ret+8(FP)
	RET
