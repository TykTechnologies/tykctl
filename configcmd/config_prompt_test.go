package configcmd

import "testing"

func TestNameMatch(t *testing.T) {
	tests := []struct {
		name string
		file string
		want bool
	}{
		{
			name: "Test Ok",
			file: "itachi-m",
			want: true,
		},
		{
			name: "Test no hypen",
			file: "itachi",
			want: true,
		},
		{
			name: "Test underscore",
			file: "itachi_m",
			want: false,
		},
		{
			name: "Test space",
			file: "itachi k",
			want: false,
		},
		{
			name: "Test extension",
			file: "itachi.yml",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nameMatch(tt.file); got != tt.want {
				t.Errorf("nameMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
