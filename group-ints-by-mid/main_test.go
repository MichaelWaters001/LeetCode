package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		list []int
	}
	wantgroup := make(map[int]int)
	wantgroup[12345]=3
	wantgroup[12346]=1
	wantgroup[23456]=1
	wantgroup[33456]=1
	wantgroup[98765]=1

	wantgroup2 := make(map[int]int)
	wantgroup2[12345]=3
	wantgroup2[12346]=1

	tests := []struct {
		name       string
		args       args
		wantGroups map[int]int
	}{
		{
			name: "test1",
			args: args{
				list: []int{18123451111, 90123452222, 28123451234, 10123461111, 46234561111, 65334561111, 90987654321}},
			wantGroups: wantgroup,
		},
		{
			name: "test2",
			args: args{
				list: []int{18123451111, 90123452222, 4444}},
			wantGroups: wantgroup,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGroups := solution(tt.args.list); !reflect.DeepEqual(gotGroups, tt.wantGroups) {
				t.Errorf("solution() = %v, want %v", gotGroups, tt.wantGroups)
			}
		})
	}
}
