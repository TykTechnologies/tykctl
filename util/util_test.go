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

func TestAbbreviateDirection(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Test North",
			input: "north",
			want:  "n",
		},
		{
			name:  "Test West",
			input: "west",
			want:  "w",
		},
		{
			name:  "Test north east",
			input: "northeast",
			want:  "ne",
		},
		{
			name:  "Test South West",
			input: "southwest",
			want:  "sw",
		},
		{
			name:  "test South East",
			input: "southeast",
			want:  "se",
		},
		{
			name:  "Test North West",
			input: "northwest",
			want:  "nw",
		},

		{
			name:  "Test Central",
			input: "central",
			want:  "c",
		},
		{
			name:  "Test South",
			input: "south",
			want:  "s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbbreviateDirection(tt.input); got != tt.want {
				t.Errorf("AbbreviateDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		elements []string
		element  string
		want     bool
	}{
		{
			name:     "Test String  does not contains",
			elements: []string{"hello", "to ", "go", "me"},
			element:  "t",
			want:     false,
		},
		{
			name:     "Test String contains",
			elements: []string{"mmh", "home ", "to", "gone"},
			element:  "gone",
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.elements, tt.element); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateUrlFromZone(t *testing.T) {
	tests := []struct {
		name    string
		region  string
		want    string
		wantErr bool
	}{
		{
			name:    "Test correct url",
			region:  "aws-eu-west-2",
			want:    "https://controller-aws-euw2.cloud-ara.tyk.io:37001",
			wantErr: false,
		},
		{
			name:    "Test wrong region format",
			region:  "aws-eu-west",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateURLFromZone(tt.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateUrlFromZone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateUrlFromZone() got = %v, want %v", got, tt.want)
			}
		})
	}
}
