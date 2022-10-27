package cmd

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

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
			err := CreateConfigFile(tt.args.dir, tt.args.file)
			assert.Equal(t, tt.Error, err)
			assert.FileExists(t, filepath.Join(tt.args.dir, tt.args.file))
			if tt.Error == nil {
				err := CreateConfigFile(tt.args.dir, tt.args.file)
				assert.Nil(t, err)
			}
			err = os.Remove(filepath.Join(tt.args.dir, tt.args.file))
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
