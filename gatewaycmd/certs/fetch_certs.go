package certs

import (
	"context"
	"errors"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

func NewFetchCertsCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Fetch).
		WithFlagAdder(false, fetchCertFlags).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())
			shared.AddGatewayServers(apimClient.Client.GetConfig())

			return cmd.MarkFlagRequired(shared.OrgID)
		}).MaximumArgs(1000, func(ctx context.Context, cmd cobra.Command, args []string) error {
		orgID, err := cmd.Flags().GetString(shared.OrgID)
		if err != nil {
			return err
		}

		certs, err := getCerts(ctx, apimClient.Client.CertsAPI, orgID)
		if err != nil {
			return err
		}

		return internal.PrintJSON(certs)
	})
}

func getCerts(ctx context.Context, api apim.CertsAPI, orgID string) (*apim.ListCerts200Response, error) {
	listCerts, resp, err := api.ListCertsExecute(api.ListCerts(ctx).OrgId(orgID))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return listCerts, nil
}

func fetchCertFlags(f *pflag.FlagSet) {
	f.StringP(shared.OrgID, "o", "", " Organisation ID to list the certificates")
}
