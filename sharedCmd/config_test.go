package sharedCmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
