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
func GetAlerts(w http.ResponseWriter, r *http.Request) {
	on400, err := try.Handle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	on500, err := try.Handle()
	if err != nil {
		http.Error(w, err.Error(), http.InternalServerError)
		return
	}

	try.Check(r.ParseForm())(on400)
	orgID := try.Check1(strconv.Atoi(r.Form.Get("orgId")))(on400)
	alerts := try.Check1(repository.FetchAlerts(orgID))(on500)
	body := try.Check1(json.Marshal(alerts))(on500)
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
n := try.Check1(os.Stdin.Read(buf))(cp, try.WithIgnore(io.EOF), try.WithDescription("failed to read"))
return buf[:n], nil
```
