package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{prices: []int{7,1,5,3,6,4}},
			want: 7,
		},
		{
			name: "example 2",
			args: args{prices: []int{1,2,3,4,5}},
			want: 4,
		},
		{
			name: "example 3",
			args: args{prices: []int{}},
			want: 0,
		},
		{
			name: "my test 1",
			args: args{prices: []int{1}},
			want: 0,
		},
		{
			name: "my test 2",
			args: args{prices: []int{1,3,5,1,1,2,5}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
