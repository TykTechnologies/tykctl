package cmd

import (
	"context"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

const orgListDesc = `
This command will list all your organizations.
Currently you can only be part of one organization hence we will return a single organization.
Sample command usage:
tykctl cloud org list --output<json/table>
You can get the output either in table or json format.The default is table format.
user the --output flag to change the format.
`

func NewOrgListCommand() *cobra.Command {
	return NewCmd("list").
		WithExample("tykctl cloud org list --output<json/table>").
		WithLongDescription(orgListDesc).
		NoArgs(func(ctx context.Context, command cobra.Command) error {
			log.Println("this is to list organization")
			getOrg(command.Context())
			return nil
		})
}

func getOrg(ctx context.Context) {
	config := &cloud.Configuration{
		DefaultHeader: map[string]string{},
		Servers: []cloud.ServerConfiguration{{
			URL:         "https://controller-aws-eun1.ara-staging.tyk.technology:37001",
			Description: "",
		},
		},
	}
	token := fmt.Sprintf("Bearer %s", viper.GetString("token"))
	config.AddDefaultHeader("Authorization", token)
	client := cloud.NewAPIClient(config)
	org, res, err := client.OrganisationsApi.GetOrgs(ctx).Execute()
	log.Println(res)
	if err != nil {
		log.Println(res.Request.URL.String())
		log.Println(res.Request.Header)
		log.Println(org.GetError())
		log.Println(err)
		return
	}
	log.Println(org.Payload)

}

func addOrgListFlags(f *pflag.FlagSet) {

}
