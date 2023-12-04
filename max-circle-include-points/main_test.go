package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		S string
		X []int
		Y []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "my 6",
			args: args{
				S: "AAABB",
				X: []int{3,2,1,2,3},
				Y: []int{0,0,0,0,0},
			},
			want: 1,
		},
		{
			name: "my 5",
			args: args{
				S: "AAA",
				X: []int{3,2,1},
				Y: []int{0,0,0},
			},
			want: 1,
		},
		{
			name: "my 4",
			args: args{
				S: "AAA",
				X: []int{1,2,3},
				Y: []int{0,0,0},
			},
			want: 1,
		},
		{
			name: "my 3",
			args: args{
				S: "ABDCAC",
				X: []int{2, -1, -4, -3, 3, -3},
				Y: []int{2, -2, 4, 1, -3, -1},
			},
			want: 2,
		},
		{
			name: "my 2",
			args: args{
				S: "ABDCAB",
				X: []int{2, -1, -4, -3, 3, 2},
				Y: []int{2, -2, 4, 1, -3, 1},
			},
			want: 0,
		},
		{
			name: "sample",
			args: args{
				S: "ABDCA",
				X: []int{2, -1, -4, -3, 3},
				Y: []int{2, -2, 4, 1, -3},
			},
			want: 3,
		},
		{
			name: "sample2",
			args: args{
				S: "ABB",
				X: []int{1, -2, -2},
				Y: []int{1, -2, 2},
			},
			want: 1,
		},
		{
			name: "sample3",
			args: args{
				S: "CCD",
				X: []int{1, -1, 2},
				Y: []int{1, -1, -2},
			},
			want: 0,
		},
		{
			name: "my1",
			args: args{
				S: "",
				X: []int{},
				Y: []int{},
			},
			want: 0,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S, tt.args.X, tt.args.Y); got != tt.want {
				t.Errorf("Solution() = %v, want %v on %s", got, tt.want, tt.name)
			}
		})
	}
}
