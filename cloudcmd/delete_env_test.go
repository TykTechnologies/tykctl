package cloudcmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/tykctl/internal"
)

func TestNewDeleteEnvCmd(t *testing.T) {
	type args struct {
		factory internal.CloudFactory
	}
	tests := []struct {
		name string
		args args
		want *cobra.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewDeleteEnvCmd(tt.args.factory), "NewDeleteEnvCmd(%v)", tt.args.factory)
		})
	}
}
