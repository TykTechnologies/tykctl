package keys

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

func NewCreateKeyCmd(apimClient internal.ApimClient) *cobra.Command {
	return internal.NewCmd(shared.Create).
		WithFlagAdder(false, createKeyFlags).
		AddPreRunFuncs(func(cmd *cobra.Command, args []string) error {
			shared.AddGatewaySecret(apimClient.Client.GetConfig())
			shared.AddGatewayServers(apimClient.Client.GetConfig())
			return cmd.MarkFlagRequired(file)
		}).NoArgs(func(ctx context.Context, cmd cobra.Command) error {
		filePath, err := cmd.Flags().GetString(file)
		if err != nil {
			return err
		}

		sets, err := cmd.Flags().GetStringArray(shared.SetFlag)
		if err != nil {
			return err
		}

		fileReader := internal.FileTemplateReader{}

		status, err := readTemplateAndCreateKey(cmd.Context(), apimClient.Client.KeysAPI, fileReader, filePath, sets)
		if err != nil {
			return err
		}

		return internal.PrintJSON(status)
	})
}

func createKeyFlags(f *pflag.FlagSet) {
	f.StringP(file, "f", "", "the path to the file that has the api definition")
	f.StringArrayP(shared.SetFlag, "s", nil, "Set template API definition field value")
}

func readTemplateAndCreateKey(ctx context.Context, api apim.KeysAPI, reader internal.TemplateReader, name string, sets []string) (*apim.ApiModifyKeySuccess, error) {
	s, err := internal.ReadTemplateAndSetValues(reader, name, sets)
	if err != nil {
		return nil, err
	}

	var sessionState apim.SessionState

	err = json.Unmarshal([]byte(s), &sessionState)
	if err != nil {
		return nil, err
	}

	return createKey(ctx, api, sessionState)
}

func createKey(ctx context.Context, api apim.KeysAPI, state apim.SessionState) (*apim.ApiModifyKeySuccess, error) {
	status, resp, err := api.AddKeyExecute(api.AddKey(ctx).SessionState(state))
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return status, nil
}
