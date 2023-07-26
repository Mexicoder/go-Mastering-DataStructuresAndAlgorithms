package main

import (
	"reflect"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				arr1: []int{0, 3, 4, 31},
				arr2: []int{4, 6, 30},
			},
			want: []int{0, 3, 4, 4, 6, 30, 31},
		},
		{
			args: args{
				arr1: []int{0, 3, 4, 31},
				arr2: []int{},
			},
			want: []int{0, 3, 4, 31},
		},
		{
			args: args{
				arr1: []int{},
				arr2: []int{0, 3, 4, 31},
			},
			want: []int{0, 3, 4, 31},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortedArrays(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
