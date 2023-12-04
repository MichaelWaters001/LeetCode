package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{
				A: 12,
			},
			want: 12,
		},
		{
			name: "sample1",
			args: args{
				A: 123456,
			},
			want: 162534,
		},
		{
			name: "sample2",
			args: args{
				A: 130,
			},
			want: 103,
		},
		{
			name: "mine",
			args: args{
				A: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A, ); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
