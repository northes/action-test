package main

import (
	"testing"
)

func TestIsEven(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{0}, true},
		{"1", args{1}, false},
		{"2", args{2}, true},
		{"10", args{10}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.i); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
