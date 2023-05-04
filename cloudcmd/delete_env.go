package cloudcmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
)

func NewDeleteEnvCmd(factory internal.CloudFactory) *cobra.Command {
	envObject := EnvObject{}

	return NewDeleteBaseCmd(factory, &envObject, Env)
}

type EnvObject struct {
	EnvResponse *cloud.Loadout
}

func (e *EnvObject) Delete(ctx context.Context, client internal.CloudClient, config internal.UserConfig, id string, f *pflag.FlagSet) error {
	envResponse, err := validateFlagsAndDeleteEnv(ctx, client, config, id, f)
	if err != nil {
		return err
	}

	e.EnvResponse = envResponse

	return nil
}

func (e *EnvObject) GetUID() string {
	if e.EnvResponse == nil {
		return ""
	}

	return e.EnvResponse.UID
}

func validateFlagsAndDeleteEnv(ctx context.Context, client internal.CloudClient, config internal.UserConfig, envID string, f *pflag.FlagSet) (*cloud.Loadout, error) {
	envFlags, err := validateCommonEnvFlags(config)
	if err != nil {
		return nil, err
	}

	cascade, err := f.GetBool(cascade)
	if err != nil {
		return nil, err
	}

	return deleteEnv(ctx, client, envFlags.OrgID, envFlags.TeamID, envID, cascade)
}

func deleteEnv(ctx context.Context, client internal.CloudClient, orgID, teamID, id string, cascade bool) (*cloud.Loadout, error) {
	localVarOptionals := cloud.LoadoutsApiDeleteLoadoutOpts{}

	if cascade {
		localVarOptionals.Cascade = optional.NewString("cascade")
	}

	deleteResponse, resp, err := client.DeleteEnv(ctx, orgID, teamID, id, &localVarOptionals)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		return nil, fmt.Errorf("returned error code %d while deleting env", resp.StatusCode)
	}

	return deleteResponse.Payload, nil
}

func shouldDelete(prompt internal.CloudPrompt, object string, f *pflag.FlagSet) (bool, error) {
	confirmed, err := f.GetBool(confirm)
	if err != nil {
		return false, err
	}

	if confirmed {
		return true, nil
	}

	return prompt.PerformActionPrompt(object)
}

type CloudObjectType int64

const (
	Org CloudObjectType = iota
	Team
	Env
	Dep
)
