package main

import "errors"

var (
	errSessionInvalid = errors.New("session invalid")
	errNoThread       = errors.New("thread not exists")
)
