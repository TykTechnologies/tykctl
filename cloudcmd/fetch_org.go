package cloudcmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
)

const fetchOrgDesc = `
This command will fetch your organizations.

Currently you can only be part of one organization hence we will return a single organization.

Sample command usage:

tykctl cloud orgs fetch --output<json/table>

You can get the output either in table or json format.The default is table format.
user the --output flag to change the format.
`

var (
	ErrorFetchingOrg  = errors.New("error fetching organization")
	ErrorGenericError = errors.New("error fetching data")
	ErrorOutPutFormat = errors.New("output flag only allows table or json")
)

func NewOrgListCommand(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(fetch).
		WithExample("tykctl cloud orgs fetch --output<json/table>").
		WithFlagAdder(false, addOutPutFlags).
		WithLongDescription(fetchOrgDesc).
		WithDescription("fetch the organizations you belong to.").
		AddPreRunFuncs(NewCloudRbac(OrgAdmin, factory.Config).CloudRbac).
		MaximumArgs(1, func(ctx context.Context, cmd cobra.Command, args []string) error {
			outPut, err := cmd.Flags().GetString(outPut)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			if len(args) == 1 {
				err = FetchAndPrintOrgByID(cmd.Context(), factory.Client, outPut, args[0])
				if err != nil {
					cmd.PrintErrln(err)
					return err
				}
				return nil
			}
			err = FetchAndPrintOrganizations(ctx, factory.Client, outPut)
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			return nil
		})
}

// FetchAndPrintOrganizations send a prints our organizations either as json or as a table.
func FetchAndPrintOrganizations(ctx context.Context, client internal.CloudClient, output string) error {
	if output != table && output != jsonFormat {
		return ErrorOutPutFormat
	}

	org, err := GetOrgs(ctx, client)
	if err != nil {
		return err
	}

	if output == table {
		internal.Printable(CreateOrgHeaderAndRows(org))
		return nil
	}

	return internal.PrintJSON(org)
}

// FetchAndPrintOrgByID send a prints a single organization either as json or as a table.
func FetchAndPrintOrgByID(ctx context.Context, client internal.CloudClient, output, oid string) error {
	if output != table && output != jsonFormat {
		return ErrorOutPutFormat
	}

	organization, err := GetOrgByID(ctx, client, oid)
	if err != nil {
		return err
	}

	if output == table {
		var organizations []cloud.Organisation
		if organization != nil {
			organizations = append(organizations, *organization)
		}

		internal.Printable(CreateOrgHeaderAndRows(organizations))

		return nil
	}

	return internal.PrintJSON(organization)
}

// GetOrgByID fetches a single organisation using its uuid.
func GetOrgByID(ctx context.Context, client internal.CloudClient, oid string) (*cloud.Organisation, error) {
	orgResponse, resp, err := client.GetOrgByID(ctx, oid)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrorFetchingOrg
	}

	if orgResponse.Status != statusOK {
		return nil, errors.New(orgResponse.Error_)
	}

	return orgResponse.Payload, nil
}

// GetOrgs fetches all organisation that belongs to a user.
func GetOrgs(ctx context.Context, client internal.CloudClient) ([]cloud.Organisation, error) {
	orgResponse, resp, err := client.GetOrgs(ctx)
	if err != nil {
		return nil, errors.New(internal.ExtractErrorMessage(err))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrorFetchingOrg
	}

	if orgResponse.Status != statusOK {
		return nil, errors.New(orgResponse.Error_)
	}

	if orgResponse.Payload == nil {
		return nil, nil
	}

	return orgResponse.Payload.Organisations, nil
}

// CreateOrgHeaderAndRows will take a list of organization and return headers and rows to be used to draw tables.
func CreateOrgHeaderAndRows(organizations []cloud.Organisation) ([]string, [][]string) {
	header := []string{"Name", "ID", "Teams", "Environments", "Control planes", "Edge"}
	rows := make([][]string, 0)

	for _, organization := range organizations {
		row := []string{
			organization.Name, organization.UID,
			getEntitlements(organization.Entitlements.Counters, teamEntitlement),
			getEntitlements(organization.Entitlements.Counters, environmentEntitlement),
			getEntitlements(organization.Entitlements.Counters, dashboardEntitlement),
			getEntitlements(organization.Entitlements.Counters, gatewayEntitlement),
		}
		rows = append(rows, row)
	}

	return header, rows
}

func getEntitlements(counter map[string]cloud.CounterEntitlement, key string) string {
	if counterEntitlement, ok := counter[key]; ok {
		return fmt.Sprintf("%d of %d", counterEntitlement.Consumed, counterEntitlement.Allowed)
	}

	return "- of -"
}
