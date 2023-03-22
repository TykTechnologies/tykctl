package templates

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/tidwall/sjson"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/gatewaycmd/shared"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
)

const (
	fromTemplate = "from-template"
	fromAPi      = "from-api"
	tempDIR      = "template-dir"
)

func NewCreateTemplate() *cobra.Command {
	return internal.NewCmd(shared.Create).
		WithFlagAdder(false, templateFlags).
		ExactArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			templateDIR, err := cmd.Flags().GetString(tempDIR)
			if err != nil {
				return err
			}

			fromTemp, err := cmd.Flags().GetString(fromTemplate)
			if err != nil {
				return err
			}

			sets, err := cmd.Flags().GetStringArray(shared.SetFlag)
			if err != nil {
				return err
			}

			err = createFromTemplate(args[0], templateDIR, fromTemp, sets)
			if err != nil {
				return err
			}

			fmt.Printf("Successfully created template \"%s\"\n", args[0])

			return nil
		})
}

func createFromTemplate(name, dir, template string, sets []string) error {
	err := util.CheckDirectory(dir)
	if err != nil {
		return err
	}

	fileName := name + ".json"
	path := filepath.Join(dir, fileName)

	_, err = os.Stat(path)
	if err == nil {
		return fmt.Errorf("template %s already exists", name)
	}

	if !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	var api *apim.APIDefinition
	if template == shared.HTTPBin {
		api = createLeanKeylessApiDefinition()
	}

	api, err = createApi(name, api, sets)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	bytes, err := json.MarshalIndent(api, "", "  ")
	if err != nil {
		return err
	}
	_, err = f.Write(bytes)

	return err
}

func templateFlags(f *pflag.FlagSet) {
	f.StringP(fromTemplate, "t", "httpbin", "Templates to use to create the template file")
	f.StringP(tempDIR, "d", ".", "where to store the template")
	f.StringArrayP(shared.SetFlag, "s", nil, "Set template API definition field value")
}

func createApi(name string, api *apim.APIDefinition, sets []string) (*apim.APIDefinition, error) {
	if api == nil {
		api = new(apim.APIDefinition)
	}
	api.Name = util.GetStrPtr(name)
	api.Slug = util.GetStrPtr(name)
	if api.Proxy == nil {
		api.Proxy = new(apim.APIDefinitionProxy)
	}
	api.Proxy.ListenPath = util.GetStrPtr(name)
	apiJson, err := json.Marshal(api)
	if err != nil {
		return nil, err
	}

	j := string(apiJson)

	for _, set := range sets {
		keyValue := strings.Split(set, "=")
		if keyValue[1] == "true" || keyValue[1] == "false" {
			value, err := strconv.ParseBool(keyValue[1])
			if err != nil {
				return nil, err
			}
			j, err = sjson.Set(j, keyValue[0], value)
		} else {
			j, err = sjson.Set(j, keyValue[0], keyValue[1])
		}

		if err != nil {
			return nil, err
		}
	}
	err = json.Unmarshal([]byte(j), &api)
	if nil != err {
		return nil, err
	}

	return api, nil
}
