package main

import (
	"testing"

	"go.uber.org/goleak"
)

// go test -v
func TestSucc(t *testing.T) {
	defer goleak.VerifyNone(t)
	Succ()
}

func TestFail(t *testing.T) {
	defer goleak.VerifyNone(t)
	Fail()
}
