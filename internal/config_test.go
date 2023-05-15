package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllConfig(t *testing.T) {
	assert.FileExists(t, "./testdata/overwritten.yml")
	assert.FileExists(t, "./testdata/config_default.yaml")
	assert.FileExists(t, "./testdata/config_itachi.yml")
	assert.FileExists(t, "./testdata/config_json.json")

	configFiles, err := GetAllConfig("./testdata")
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"config_default.yaml", "config_itachi.yml"}, configFiles)
}

func TestConfigFileNotOverWritten(t *testing.T) {
	assert.FileExists(t, "./testdata/overwritten.yml")

	err := CreateFile("./testdata", "overwritten.yml")
	assert.Nil(t, err)

	content, err := os.ReadFile("./testdata/overwritten.yml")
	assert.Nil(t, err)
	assert.Equal(t, "data: testdata", string(content))
}

func TestCreateConfigFile(t *testing.T) {
	type args struct {
		dir  string
		file string
	}

	tests := []struct {
		name  string
		args  args
		Error error
	}{
		{
			name: "Test File created",
			args: args{
				dir:  "./testdata",
				file: "config.yml",
			},
			Error: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateFile(tt.args.dir, tt.args.file)
			assert.Equal(t, tt.Error, err)
			assert.FileExists(t, filepath.Join(tt.args.dir, tt.args.file))
			err = os.Remove(filepath.Join(tt.args.dir, tt.args.file))
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

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
