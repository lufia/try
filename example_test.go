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
		// Output: error: unsupported operation
		return
	}
	v := try.Check(read())(s)
	fmt.Println("value:", v)
}

func read() (string, error) {
	return "", errors.ErrUnsupported
}
