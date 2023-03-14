package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateKeyFromPath(t *testing.T) {
	tests := []struct {
		name  string
		paths []string
		want  string
	}{
		{
			name:  "Test empty path",
			paths: []string{},
			want:  "",
		},
		{
			name:  "Test single path",
			paths: []string{"cloud"},
			want:  "cloud",
		},
		{
			name:  "Test two paths",
			paths: []string{"cloud", "team"},
			want:  "cloud.team",
		},
		{
			name:  "Test three paths",
			paths: []string{"cloud", "team", "name"},
			want:  "cloud.team.name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CreateKeyFromPath(tt.paths...))
		})
	}
}
