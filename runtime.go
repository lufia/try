package try

type stack struct {
	lo uintptr
	hi uintptr
}

type g struct {
	stack stack
}
