package main

import "testing"

func TestHashTable_hash(t *testing.T) {
	type fields struct {
		size int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "grapes",
			fields: fields{
				size: 50,
			},
			args: args{
				key: "grapes",
			},
			want: 23,
		},
		{
			name: "apples",
			fields: fields{
				size: 50,
			},
			args: args{
				key: "apples",
			},
			want: 39,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashTable(tt.fields.size)
			if got := h.hash(tt.args.key); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}

}

//func TestHashTable_hashCollision(t *testing.T) {
//	hashTarget := 22
//	generatedHash := 0
//	h := NewHashTable(50)
//	for {
//
//		hash := h.hash()
//	}
//}
