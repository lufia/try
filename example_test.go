package try_test

import (
	"errors"
	"fmt"

	"github.com/lufia/try"
)

func Example() {
	s, err := try.Handle()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	v := try.Check1(read())(s)
	fmt.Println("value:", v)
	// Output:
	// error: unsupported operation
}

func read() (string, error) {
	return "", errors.ErrUnsupported
}
