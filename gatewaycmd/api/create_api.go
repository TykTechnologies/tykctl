package api

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

func NewCreateAPICmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Create).
		WithFlagAdder(false, createAPIFlags).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())

			err := shared.AddGatewayServers(apimClient.Client.GetConfig())
			if err != nil {
				return err
			}

			return cmd.MarkFlagRequired(file)
		}).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			filePath, err := cmd.Flags().GetString(file)
			if err != nil {
				return err
			}

			sets, err := cmd.Flags().GetStringArray(shared.SetFlag)
			if err != nil {
				return err
			}

			fileReader := internal.FileTemplateReader{}
			status, err := readTemplateAndCreateAPI(cmd.Context(), apimClient.Client.APIsAPI, fileReader, filePath, sets)
			if err != nil {
				return err
			}

			return internal.PrintJSON(status)
		})
}

func readTemplateAndCreateAPI(ctx context.Context, api apim.APIsAPI, reader internal.TemplateReader, name string, sets []string) (*apim.ApiModifyKeySuccess, error) {
	s, err := internal.ReadTemplateAndSetValues(reader, name, sets)
	if err != nil {
		return nil, err
	}

	var definition apim.APIDefinition

	err = json.Unmarshal([]byte(s), &definition)
	if err != nil {
		return nil, err
	}

	return createAPI(ctx, api, definition)
}

func createAPIFlags(f *pflag.FlagSet) {
	f.StringP(file, "f", "", "the path to the file that has the api definition")
	f.StringArrayP(shared.SetFlag, "s", nil, "Set template API definition field value")
}

func createAPI(ctx context.Context, api apim.APIsAPI, definition apim.APIDefinition) (*apim.ApiModifyKeySuccess, error) {
	status, resp, err := api.CreateApiExecute(api.CreateApi(ctx).APIDefinition(definition))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return status, nil
}
