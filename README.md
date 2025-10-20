# try
An experimental error handling library.

[![GoDev][godev-image]][godev-url]
[![Actions Status][actions-image]][actions-url]

## Supported Architectures

* amd64
* arm64

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

*try.Handle* creates a fallback point, called "checkpoint",  then return nil error
 at the first time.

After that, above code calls *os.ReadFile* and *url.Parse* with *try.Check*. If either these functions returns an error, *try.Check* rewind the program to the checkpoint, then *try.Handle* will return the error.

**I strongly recommend that Check and Handle should call on the same stack.**

## Example

Error handling in Go sometimes gets flustrated. For instance:

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
	cp400, err := try.Handle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cp500, err := try.Handle()
	if err != nil {
		http.Error(w, err.Error(), http.InternalServerError)
		return
	}

	try.Raise(r.ParseForm())(cp400)
	orgID := try.Check(strconv.Atoi(r.Form.Get("orgId")))(cp400)
	alerts := try.Check(repository.FetchAlerts(orgID))(cp500)
	body := try.Check(json.Marshal(alerts))(cp500)
	...
}
```

[godev-image]: https://pkg.go.dev/badge/github.com/lufia/try
[godev-url]: https://pkg.go.dev/github.com/lufia/try
[actions-image]: https://github.com/lufia/try/actions/workflows/test.yml/badge.svg
[actions-url]: https://github.com/lufia/try/actions/workflows/test.yml
