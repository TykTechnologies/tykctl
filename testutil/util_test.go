package testutil

import (
	"errors"
	"testing"
)

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
		{
			name: "Test equal",
			args: args{
				a: errors.New("this errors are equal"),
				b: errors.New("this errors are equal"),
			},
			want: true,
		},
		{
			name: "Test nil errors",
			args: args{
				a: nil,
				b: nil,
			},
			want: true,
		},
		{
			name: "Test nil and full error",
			args: args{
				a: errors.New("this errors are equal"),
				b: nil,
			},
			want: false,
		},

		{
			name: "Test different messages",
			args: args{
				a: errors.New("this errors are equal"),
				b: errors.New("this errors are "),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualError(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("EqualError() = %v, want %v", got, tt.want)
			}
		})
	}
}
