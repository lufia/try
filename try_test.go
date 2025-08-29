package try

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
	"unsafe"

	"github.com/m-mizutani/gt"
)

func TestHandle(t *testing.T) {
	var i int
	s, err := Handle()
	gt.NoError(t, err)

	f := runtime.FuncForPC(uintptr(s.pc))
	t.Logf("func = %s\n", f.Name())
	t.Logf("pc = %x\n", uintptr(s.pc))
	t.Logf("sp = %x\n", uintptr(s.sp))
	t.Logf("frame = %x\n", uintptr(unsafe.Pointer(&i)))
	gt.Number(t, int64(uintptr(s.pc))).NotEqual(0)
	gt.Number(t, int64(uintptr(s.sp))).NotEqual(0)
	gt.NoError(t, s.err)
}

func TestScopeRaise(t *testing.T) {
	raised := false
	s, err := Handle()
	if err != nil {
		raised = true
	}
	if !raised {
		s.Raise(errors.New("fake"))
	}
	gt.Bool(t, raised).True()
}

func TestCheck_onError(t *testing.T) {
	raised := false
	s, err := Handle()
	if err != nil {
		raised = true
	}
	if !raised {
		Check(10, errors.New("fake")).Eval(s)
	}
	gt.Bool(t, raised).True()
}

func TestCheck_onErrorWithWrap(t *testing.T) {
	wrap := func(err error) error {
		return fmt.Errorf("failed: %w", err)
	}
	msg := ""
	s, err := Handle()
	if err != nil {
		msg = err.Error()
	}
	if msg == "" {
		Check(10, errors.New("fake")).Wrap(wrap).Eval(s)
	}
	gt.String(t, msg).Equal("failed: fak")
}
