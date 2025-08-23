#!/bin/bash

# func Handle must be inlined.
go build -gcflags=-m 2>&1 | grep 'can inline Handle'
