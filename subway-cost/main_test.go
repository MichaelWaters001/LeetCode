package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		start []int
		dest  []int
		limit []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{
				start: []int{1, 0, 2, 4},
				dest:  []int{2, 2, 0, 5},
				limit: []int{3, 17, 7, 4, 5, 17},
			},
			want: 16,
		},
		{
			name: "sample2",
			args: args{
				start: []int{1,2,0,2,3},
				dest:  []int{2,1,2,1,2},
				limit: []int{4,8,18,16,20},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.start, tt.args.dest, tt.args.limit); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
