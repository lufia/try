//go:build go1.27

package try

import (
	"errors"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestCheckpoint_Check(t *testing.T) {
	msg := ""
	cp, err := Handle()
	if err != nil {
		msg = err.Error()
	}
	if msg == "" {
		cp.Check(errors.New("fake"))
	}
	gt.String(t, msg).Equal("fake")
}

func TestCheckpoint_CheckOptions(t *testing.T) {
	msg := ""
	cp, err := Handle()
	if err != nil {
		msg = err.Error()
	}
	if msg == "" {
		cp.CheckOptions(errors.New("fake"))(WithDescription("failed"))
	}
	gt.String(t, msg).Equal("failed: fake")
}
