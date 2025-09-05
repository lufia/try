package main

import (
	"fmt"
	"unsafe"

	"github.com/lufia/try"
)

func getsp() uintptr
func getmem(addr uintptr) uintptr

// $104-0
func main() {
	_, err := try.Handle()
	if err != nil {
		panic(err)
	}
	var xs uintptr        // This will be set SP with delve
	var sp unsafe.Pointer // I should break here and set xs
	sp = unsafe.Pointer(xs)
	fmt.Printf("main: sp=0x%x input=0x%x\n", sp, xs)
	cook(sp)
}

// $184-8
func cook(parentSP unsafe.Pointer) {
	//targetBP := uintptr(parentSP) - 8
	targetBP := uintptr(parentSP)
	sp := getsp()
	fmt.Printf("cook: sp=0x%x\n", sp)
	for addr := sp; addr <= targetBP; addr += 8 {
		if addr == uintptr(unsafe.Pointer(&targetBP)) {
			continue
		}
		v := getmem(addr)
		if v == targetBP {
			fmt.Printf("found: target=0x%x parent=0x%x offset=0x%x\n", targetBP, parentSP, addr-sp)
			break
		}
	}
}
