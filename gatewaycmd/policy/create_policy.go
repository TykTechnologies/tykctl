package policy

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
)

const (
	file = "file"
)

func NewCreatePolicyCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Create).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())
			err := shared.AddGatewayServers(apimClient.Client.GetConfig())
			if err != nil {
				return err
			}

			return cmd.MarkFlagRequired(file)
		}).
		WithFlagAdder(false, createPolicyFlags).NoArgs(func(ctx context.Context, cmd cobra.Command) error {
		filePath, err := cmd.Flags().GetString(file)
		if err != nil {
			return err
		}

		sets, err := cmd.Flags().GetStringArray(shared.SetFlag)
		if err != nil {
			return err
		}

		fileReader := internal.FileTemplateReader{}

		status, err := readTemplateAndCreatePolicy(cmd.Context(), apimClient.Client.PoliciesAPI, fileReader, filePath, sets)
		if err != nil {
			return err
		}

		return internal.PrintJSON(status)
	})
}

func readTemplateAndCreatePolicy(ctx context.Context, api apim.PoliciesAPI, reader internal.TemplateReader, name string, sets []string) (*apim.ApiModifyKeySuccess, error) {
	s, err := internal.ReadTemplateAndSetValues(reader, name, sets)
	if err != nil {
		return nil, err
	}

	var policy apim.Policy

	err = json.Unmarshal([]byte(s), &policy)
	if err != nil {
		return nil, err
	}

	return createPolicy(ctx, api, policy)
}

func createPolicy(ctx context.Context, api apim.PoliciesAPI, policy apim.Policy) (*apim.ApiModifyKeySuccess, error) {
	status, resp, err := api.AddPolicyExecute(api.AddPolicy(ctx).Policy(policy))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return status, nil
}

func createPolicyFlags(f *pflag.FlagSet) {
	f.StringP(file, "f", "", "the path to the file that has the api definition")
	f.StringArrayP(shared.SetFlag, "s", nil, "Set template API definition field value")
}
