package util

import "testing"

func TestStringIsEmpty(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringIsEmpty(tt.args.v); got != tt.want {
				t.Errorf("StringIsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
