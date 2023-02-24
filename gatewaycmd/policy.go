package gatewaycmd

import (
	"context"
	"errors"
	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
	"net/http"
)

// NewPolicyCommand policy create policy.
func NewPolicyCommand() *cobra.Command {
	return internal.NewCmd("policy").NoArgs(func(ctx context.Context, cmd cobra.Command) error {
		cmd.Println("hello I am policy")
		return nil
	})

}

func getPolicy(client apim.APIClient) ([]apim.Policy, error) {
	policies, h, err := client.PoliciesAPI.ListPolicies(context.Background()).Execute()
	if err != nil {
		return nil, err
	}
	if h.StatusCode != http.StatusOK {
		return nil, errors.New(h.Status)
	}
	return policies, nil

}
