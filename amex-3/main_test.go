package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				A: []int{1, 3, 6, 4, 1, 2},
			},
			want: 5,
		},
		{
			name: "test2",
			args: args{
				A: []int{1, 2, 3},
			},
			want: 4,
		},
		{
			name: "test3",
			args: args{
				A: []int{-1, -3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
