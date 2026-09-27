# try

An experimental error handling library.

[![GoDev][godev-image]][godev-url]
[![Actions Status][actions-image]][actions-url]

This library, **try**, aims to reduce `if err != nil`, and its design is inspired by [Russ's Error Handling proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling-overview.md).

[godev-image]: https://pkg.go.dev/badge/github.com/lufia/try
[godev-url]: https://pkg.go.dev/github.com/lufia/try
[actions-image]: https://github.com/lufia/try/actions/workflows/test.yml/badge.svg
[actions-url]: https://github.com/lufia/try/actions/workflows/test.yml

## Supported Architectures

For now it supports only three architectures:

* amd64
* arm64
* 386

## Usage

This is a simple example.

```go
import (
	"net/url"
	"os"

	"github.com/lufia/try"
)

func Run(file string) (string, error) {
	cp, err := try.Handle()
	if err != nil {
		return "", err
	}
	s := cp.Check1(os.ReadFile(file))
	u := cp.Check1(url.Parse(string(s)))

	// If you use Go 1.26 or earlier, use function style instead.
	s := try.Check1(os.ReadFile(file))(cp)
	u := try.Check1(url.Parse(string(s)))(cp)

	return u.Path, nil
}
```

*try.Handle* creates a fallback point, called "checkpoint", then returns a nil error the first time.

After that, the above code calls *os.ReadFile* and *url.Parse* with *try.Check1*. If either of these functions returns an error, *try.Check1* rewinds the program to the checkpoint, and *try.Handle* then returns the error.

**I strongly recommend calling *Check* and *Handle* on the same stack.**

## Example

Error handling in Go sometimes makes developers frustrated. For instance:

```go
func GetAlerts(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	orgID, err := strconv.Atoi(r.Form.Get("orgId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	alerts, err := repository.FetchAlerts(orgID)
	if err != nil {
		http.Error(w, err.Error(), http.InternalServerError)
		return
	}
	body, err := json.Marshal(alerts)
	if err != nil {
		http.Error(w, err.Error(), http.InternalServerError)
		return
	}
	...
}
```

The example above can rewrite more simple with **try**.

```go
type httpError struct {
	Err  error
	Code int
}

func (e *httpError) Error() string { return e.Err.Error() }

func (e *httpError) Unwrap() error { return e.Err }

func e(err error, code int) *httpError {
	if err == nil { return nil }
	return &httpError{err, code}
}

func GetAlerts(w http.ResponseWriter, r *http.Request) {
	cp, err := try.HandleFor[*httpError]()
	if err != nil {
		http.Error(w, http.StatusText(err.Code), err.Code)
		return
	}
	cp.Check(e(r.ParseForm(), http.StatusBadRequest))

	orgID, err := strconv.Atoi(r.Form.Get("orgId"))
	cp.Check(e(err, http.StatusBadRequest))

	alerts, err := repository.FetchAlerts(orgID)
	cp.Check(e(err, http.StatusInternalServerError))

	body, err := json.Marshal(alerts)
	cp.Check(e(err, http.StatusInternalServerError))
	...
}
```

### Options

There are three options for *Check* and its variants.

```go
cp, err := try.Handle()
if err != nil {
	return nil, err
}
buf := make([]byte, 1<<8)

n := cp.Check1Options(os.Stdin.Read(buf))(
	try.WithIgnore(io.EOF),
	try.WithDescription("failed to read"),
)

// If you use Go 1.26 or earlier, use function style instead.
n := try.Check1(os.Stdin.Read(buf))(cp, try.WithIgnore(io.EOF), try.WithDescription("failed to read"))

return buf[:n], nil
```
