package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "test 1",
			args: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args)
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", tt.args, tt.want)
			}
			fmt.Println(tt.args)
		})
	}
}
