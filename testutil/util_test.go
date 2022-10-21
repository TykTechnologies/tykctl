package testutil

import "testing"

func TestEqualError(t *testing.T) {
	type args struct {
		a error
		b error
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
			if got := EqualError(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("EqualError() = %v, want %v", got, tt.want)
			}
		})
	}
}
