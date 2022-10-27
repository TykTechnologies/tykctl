package cmd

import (
	"context"
	"github.com/TykTechnologies/tykctl/internal"
	"testing"
)

func TestGetOrg(t *testing.T) {
	type args struct {
		ctx    context.Context
		client internal.CloudClient
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetOrg(tt.args.ctx, tt.args.client)
		})
	}
}
