# try
An experimental error handling library.

## Supported Architectures

* amd64

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
	u := try.Check(url.Parse(s))(scope)
	return u.Path, nil
}
```
