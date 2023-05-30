package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringCompression(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 1",
			args{s: "aabbb"},
			"a2b3",
		},
		{
			"Test 2",
			args{s: "ab"},
			"ab",
		},
		{
			"Test 3",
			args{s: "aabb"},
			"aabb",
		},
		{
			"Test 4",
			args{s: "aabcccccaaa"},
			"a2b1c5a3",
		},
		{
			"Test 5",
			args{s: "aaa"},
			"a3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StringCompression(tt.args.s), "StringCompression(%v)", tt.args.s)
		})
	}
}
