package main

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1123",
			args: args{	s: "1123"},
			want: 5,
		},
		{
			name: "10",
			args: args{	s: "10"},
			want: 1,
		},
		{
			name: "2101",
			args: args{	s: "2101"},
			want: 1,
		},
		{
			name: "1012",
			args: args{	s: "1012"},
			want: 2,
		},	
		{
			name: "123456789",
			args: args{	s: "123456789"},
			want: 3,
		},
		{
			name: "226",
			args: args{	s: "226"},
			want: 3,
		},
		{
			name: "12",
			args: args{	s: "12"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings(%s) = %v, want %v",tt.name, got, tt.want)
			}
		})
	}
}
