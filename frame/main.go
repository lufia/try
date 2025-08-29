package main

import (
	"errors"
	"fmt"
)

func getsp() uintptr
func getfp() uintptr
func returnTo(off int) uintptr
func returnTo2(off int) uintptr

func main() {
	sp := getsp()
	fp := getfp()
	for i := range 10 {
		addr := returnTo(i * 8)
		fmt.Printf("main[%d]: sp=0x%x fp=0x%x addr=0x%x\n", i*8, sp, fp, addr)
	}
	cook()
}

func cook() {
	sp := getsp()
	fp := getfp()
	for i := range 10 {
		addr := returnTo(i * 8)
		fmt.Printf("cook[%d]: sp=0x%x fp=0x%x addr=0x%x\n", i*8, sp, fp, addr)
	}
	for i := range 5 {
		addr := returnTo2(i * 8)
		fmt.Printf("cook2[%d]: sp=0x%x fp=0x%x addr=0x%x\n", i*8, sp, fp, addr)
	}
	for i := range 5 {
		addr := returnTo2(-i * 8)
		fmt.Printf("cook2[%d]: sp=0x%x fp=0x%x addr=0x%x\n", -i*8, sp, fp, addr)
	}
}

func retret() error {
	return errors.ErrUnsupported
}
