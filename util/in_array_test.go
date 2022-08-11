package util

import "testing"

func TestInArrayString(t *testing.T) {
	type argsString struct {
		val string
		arr []string
	}

	tests := []struct {
		name       string
		args       argsString
		wantExists bool
		wantIndex  int
	}{
		{
			name: "exists",
			args: argsString{
				val: "b",
				arr: []string{"a", "b", "c", "d"},
			},
			wantExists: true,
			wantIndex:  1,
		},
		{
			name: "not exists",
			args: argsString{
				val: "e",
				arr: []string{"a", "b", "c", "d"},
			},
			wantExists: false,
			wantIndex:  -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExists, gotIndex := InArray(tt.args.val, tt.args.arr)
			if gotExists != tt.wantExists {
				t.Errorf("InArray() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("InArray() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestInArrayInt(t *testing.T) {
	type argsInt struct {
		val int
		arr []int
	}

	tests := []struct {
		name       string
		args       argsInt
		wantExists bool
		wantIndex  int
	}{
		{
			name: "exists",
			args: argsInt{
				val: 1,
				arr: []int{1, 2, 3, 4, 5},
			},
			wantExists: true,
			wantIndex:  0,
		},
		{
			name: "not exists",
			args: argsInt{
				val: 10,
				arr: []int{1, 2, 3, 4, 5},
			},
			wantExists: false,
			wantIndex:  -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExists, gotIndex := InArray(tt.args.val, tt.args.arr)
			if gotExists != tt.wantExists {
				t.Errorf("InArray() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("InArray() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}
