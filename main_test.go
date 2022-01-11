package main

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	//i := r.Int()
	//t.Log(i)
	err := Generate(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("ok")
}
