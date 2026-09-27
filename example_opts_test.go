package try_test

import (
	"fmt"
	"os"

	"github.com/lufia/try"
)

func ExampleCheck1_withOptions() {
	cp, err := try.Handle()
	if err != nil {
		fmt.Println(err)
		return
	}
	v := try.Check1(os.ReadFile("NOFILE"))(cp, try.WithDescription("failed to read a file"))
	fmt.Println("value:", v)
	// Output:
	// failed to read a file: open NOFILE: no such file or directory
}
