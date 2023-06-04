package example

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Matrix 3x3 values",
			args{matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			"Matrix 4x4 values",
			args{matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			}},
			[][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
		{
			"Matrix 2x2 values",
			args{matrix: [][]int{
				{1, 2},
				{3, 4},
			}},
			[][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			"Matrix 4x3 values",
			args{matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			}},
			[][]int{
				{10, 7, 4, 1},
				{11, 8, 5, 2},
				{12, 9, 6, 3},
			},
		},
		{
			"Matrix 1x1 values",
			args{matrix: [][]int{
				{1},
			}},
			[][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
