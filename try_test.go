package try

import (
	"errors"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestHandle(t *testing.T) {
	s, err := Handle()
	gt.NoError(t, err)

	gt.Number(t, int64(s.pc)).NotEqual(0)
	gt.Number(t, int64(s.sp)).NotEqual(0)
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
		Check(10, errors.New("fake"))(s)
	}
	gt.Bool(t, raised).True()
}
