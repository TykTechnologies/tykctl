package configcmd

import "testing"

func TestConfigFileDisplayName(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{
			name:     "Test Ok",
			fileName: "config_itachi.yml",
			want:     "itachi",
		},
		{
			name:     "Test yaml extension",
			fileName: "config_mn.yaml",
			want:     "mn",
		},
		{
			name:     "Test has no prefix",
			fileName: "conf_sasuke.yml",
			want:     "conf_sasuke",
		},
		{
			name:     "Test has no extension",
			fileName: "config_ks",
			want:     "ks",
		},
		{
			name:     "Test no prefix of suffix",
			fileName: "lj",
			want:     "lj",
		},
		{
			name:     "Test different extension",
			fileName: "config_lp.json",
			want:     "lp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConfigFileDisplayName(tt.fileName); got != tt.want {
				t.Errorf("ConfigFileDisplayName() = %v, want %v", got, tt.want)
			}
		})
	}
}
