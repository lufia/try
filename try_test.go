package try

import (
	"errors"
	"fmt"
	"runtime"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestBP(t *testing.T) {
	bp0 := getbp(1)
	bp1 := wrap(2)
	if bp0 != bp1 {
		t.Errorf("bp0 = 0x%x; bp1 = 0x%x\n", bp0, bp1)
	}
}

//go:noinline
func wrap(skip int) uintptr {
	return getbp(skip)
}

func TestHandle(t *testing.T) {
	s, err := Handle()
	gt.NoError(t, err)

	f := runtime.FuncForPC(s.pc)
	t.Logf("func = %s\n", f.Name())
	t.Logf("pc = 0x%x\n", s.pc)
	t.Logf("sp = 0x%x\n", s.sp)
	gt.Number(t, int64(s.pc)).NotEqual(0)
	gt.Number(t, int64(s.sp)).NotEqual(0)
	gt.NoError(t, s.err)
}

func TestScopeRaise(t *testing.T) {
	raised := false
	s, err := Handle()
	t.Logf("err = %v", err)
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
	gt.String(t, msg).Equal("failed: fake")
}
