#!/bin/bash
# require.sh - check whether Handle and HandleFor can be inlined

go build -gcflags=-m require.go 2>&1 |
awk '
/inlining call to / && /try\.Handle$/ {
	handle = 1
}
/inlining call to / && /try\.HandleFor\[/ {
	handlefor = 1
}
END {
	if(handle <= 0)
		print "cannot not inline Handle" >"/dev/stderr"
	if(handlefor <= 0)
		print "cannot not inline HandleFor[E]" >"/dev/stderr"
	exit(handle + handlefor == 2 ? 0 : 1)
}
'
