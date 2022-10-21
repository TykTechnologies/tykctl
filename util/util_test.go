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
		{
			name: "Test String Not empty",
			args: args{
				v: "string not empty",
			},
			want: false,
		},
		{
			name: "Test empty string",
			args: args{
				v: "",
			},
			want: true,
		},
		{
			name: "Test String with spaces",
			args: args{
				v: "      ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringIsEmpty(tt.args.v); got != tt.want {
				t.Errorf("StringIsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test valid email address",
			args: args{
				email: "support@tyk.io",
			},
			wantErr: false,
		},
		{
			name: "Test valid gmail email address",
			args: args{
				email: "support@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "Test invalid email address",
			args: args{
				email: "support",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateEmail(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
