package cmd

import (
	"context"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func NewOrgCreateCmd() *cobra.Command {
	return NewCmd("create").
		NoArgs(func(ctx context.Context, command cobra.Command) error {
			log.Println("running this command")
			createOrg(command.Context())
			return nil
		})

}

func createOrg(ctx context.Context) {
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
	organisation := cloud.Organisation{
		Name: "here test",
	}
	org, res, err := client.OrganisationsApi.CreateOrg(ctx).Organisation(organisation).Execute()
	log.Println(res)
	if err != nil {
		log.Println(org)
		log.Println(err)
		return
	}
	log.Println(org.Payload)

}
