//go:build ignore

package main

import "github.com/lufia/try"

func needInlines() {
	try.Handle()
	try.HandleFor[error]()
}
