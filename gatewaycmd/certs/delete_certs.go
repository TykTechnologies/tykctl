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

func NewDeleteCerts(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Delete).
		WithFlagAdder(false, deleteCertFlags).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())
			err := shared.AddGatewayServers(apimClient.Client.GetConfig())
			if err != nil {
				return err
			}

			return cmd.MarkFlagRequired(shared.OrgID)
		}).ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
		orgID, err := cmd.Flags().GetString(shared.OrgID)
		if err != nil {
			return err
		}

		status, err := deleteCerts(cmd.Context(), apimClient.Client.CertsAPI, orgID, args[0])
		if err != nil {
			return err
		}
		return internal.PrintJSON(status)
	})
}

func deleteCerts(ctx context.Context, api apim.CertsAPI, orgID string, certID string) (*apim.ApiStatusMessage, error) {
	status, resp, err := api.DeleteCertsExecute(api.DeleteCerts(ctx).CertID(certID).OrgId(orgID))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return status, nil
}

func deleteCertFlags(f *pflag.FlagSet) {
	f.StringP(shared.OrgID, "o", "", " Organisation ID to delete the certificates")
}
