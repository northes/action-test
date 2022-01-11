package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	i := r.Int()
	t.Log(i)
	err := Generate(i)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("ok")
}
