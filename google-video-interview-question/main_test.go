package google_video_interview_question

import "testing"

func TestHasPairWithSum_sortedData(t *testing.T) {
	type args struct {
		data []int
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// generate 5 test cases
		{
			name: "2 numbers that add up to the sum",
			args: args{
				data: []int{1, 2, 3, 9},
				sum:  8,
			},
			want: false,
		},
		{
			name: "2 numbers that do not add up to the sum",
			args: args{
				data: []int{1, 2, 4, 4},
				sum:  8,
			},
			want: true,
		},
		{
			name: "can't use the same number 2 times to get the sum",
			args: args{
				data: []int{1, 2, 4, 5},
				sum:  8,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPairWithSum_sortedData(tt.args.data, tt.args.sum); got != tt.want {
				t.Errorf("HasPairWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPairWithSum_notSortedData(t *testing.T) {
	type args struct {
		data []int
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// SORTED
		{
			name: "SORTED 2 numbers that DON'T add up to the sum",
			args: args{
				data: []int{1, 1, 3, 9},
				sum:  8,
			},
			want: false,
		},
		{
			name: "SORTED 2 numbers that DO add up to the sum",
			args: args{
				data: []int{1, 2, 4, 4},
				sum:  8,
			},
			want: true,
		},
		{
			name: "SORTED can't use the same number 2 times to get the sum",
			args: args{
				data: []int{1, 2, 4, 5},
				sum:  8,
			},
			want: false,
		},
		// NOT SORTED
		{
			name: "NOT SORTED 2 numbers that DON'T add up to the sum",
			args: args{
				data: []int{9, 1, 3, 1},
				sum:  8,
			},
			want: false,
		},
		{
			name: "NOT SORTED 2 numbers that DO not add up to the sum",
			args: args{
				data: []int{4, 2, 1, 4},
				sum:  8,
			},
			want: true,
		},
		{
			name: "NOT SORTED 2 numbers that DO not add up to the sum",
			args: args{
				data: []int{4, 2, 1, 6},
				sum:  8,
			},
			want: true,
		},
		{
			name: "NOT SORTED can't use the same number 2 times to get the sum",
			args: args{
				data: []int{5, 2, 4, 1},
				sum:  8,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPairWithSum_notSortedData(tt.args.data, tt.args.sum); got != tt.want {
				t.Errorf("HasPairWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsCommonItem1(t *testing.T) {
	type args struct {
		arr1 []string
		arr2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no common items found",
			args: args{
				arr1: []string{"a", "b", "c", "x"},
				arr2: []string{"z", "y", "i"},
			},
			want: false,
		},
		{
			name: "no common items found",
			args: args{
				arr1: []string{"a", "b", "c", "x"},
				arr2: []string{"z", "y", "x"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsCommonItem1(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("containsCommonItem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsCommonItem2(t *testing.T) {
	type args struct {
		arr1 []string
		arr2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no common items found",
			args: args{
				arr1: []string{"a", "b", "c", "x"},
				arr2: []string{"z", "y", "i"},
			},
			want: false,
		},
		{
			name: "no common items found",
			args: args{
				arr1: []string{"a", "b", "c", "x"},
				arr2: []string{"z", "y", "x"},
			},
			want: true,
		},
		{
			name: "arr1 is empty",
			args: args{
				arr1: []string{},
				arr2: []string{"z", "y", "x"},
			},
			want: false,
		},
		{
			name: "arr2 is empty",
			args: args{
				arr1: []string{"a", "b", "c", "x"},
				arr2: []string{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsCommonItem2(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("containsCommonItem2() = %v, want %v", got, tt.want)
			}
		})
	}
}
