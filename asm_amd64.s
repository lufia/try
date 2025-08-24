#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

// $32-0
#define Handle_size 32

//  0.. 7(SP) = return to Handle
//            = no frames for waserror
//  8..15(SP) = BP of Handle
// 16..23(SP) = return to main
// 24..55(SP) = frame size for Handle
// 56..63(SP) = BP of main

TEXT ·waserror(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), CX
	MOVQ	56(SP), AX
	SUBQ	$8, AX
	MOVQ	AX, Scope_sp(CX)
	MOVQ	16(SP), AX
	MOVQ	AX, Scope_pc(CX)
	MOVQ	$0, ret+8(FP)
	RET

TEXT ·raise(SB),(NOSPLIT|NOFRAME|WRAPPER),$0
	NO_LOCAL_POINTERS
	MOVQ	s+0(FP), CX
	MOVQ	Scope_sp(CX), SP
	MOVQ	Scope_pc(CX), AX
	MOVQ	AX, 0(SP)
	MOVQ	$1, ret+8(FP)
	RET
