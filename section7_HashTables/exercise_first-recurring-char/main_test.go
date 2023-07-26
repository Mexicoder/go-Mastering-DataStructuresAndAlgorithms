package main

import "testing"

func Test_firstRecurringChar(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{2, 5, 1, 2, 3, 5, 1, 2, 4},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				arr: []int{2, 1, 1, 2, 3, 5, 1, 2, 4},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				arr: []int{2, 3, 4, 5},
			},
			want: -1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstRecurringChar(tt.args.arr); got != tt.want {
				t.Errorf("firstRecurringChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
