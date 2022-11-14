package gatewaycmd

import (
	"context"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
)

// NewPolicyCommand policy create policy.
func NewPolicyCommand() *cobra.Command {
	return internal.NewCmd("policy").NoArgs(func(ctx context.Context, cmd cobra.Command) error {
		cmd.Println("hello I am policy")
		return nil
	})

}
