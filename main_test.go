package main

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	var err error
	if err = Generate(1); err == nil {
		t.Error("1 is ok,but return error")
	}

	if err = Generate(0); err != nil {
		t.Error("0 is not ok,but return nil")
	}
}
