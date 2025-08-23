# try
An experimental error handling library.

[![GoDev][godev-image]][godev-url]
[![Actions Status][actions-image]][actions-url]

## Supported Architectures

* amd64
* arm64

## Example

```go
import (
	"net/url"
	"os"

	"github.com/lufia/try"
)

func Run(file string) (string, error) {
	scope, err := try.Handle()
	if err != nil {
		return "", err
	}
	s := try.Check(os.ReadFile(file))(scope)
	u := try.Check(url.Parse(string(s)))(scope)
	return u.Path, nil
}
```

*try.Handle* creates a fallback point, called "scope",  then return nil error
 at the first time.

After that, above code calls *os.ReadFile* and *url.Parse* with *try.Check*. If either these functions returns an error, *try.Check* rewind the program to the scope, then *try.Handle* will return the error.

**I strongly recommend that Check and Handle should call on the same stack.**

[godev-image]: https://pkg.go.dev/badge/github.com/lufia/try
[godev-url]: https://pkg.go.dev/github.com/lufia/try
[actions-image]: https://github.com/lufia/try/actions/workflows/test.yml/badge.svg
[actions-url]: https://github.com/lufia/try/actions/workflows/test.yml
