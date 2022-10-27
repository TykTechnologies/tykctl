package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/spf13/cobra"
	"net/http"
)

const orgListDesc = `
This command will list all your organizations.
Currently you can only be part of one organization hence we will return a single organization.
Sample command usage:
tykctl cloud org list --output<json/table>
You can get the output either in table or json format.The default is table format.
user the --output flag to change the format.
`

var ErrorFetchingOrg = errors.New("error fetching organization")
var ErrorGenericError = errors.New("error fetching data")

const (
	teamEntitlement        = "MaxTeamCount"
	environmentEntitlement = "MaxLoadoutCount"
	dashboardEntitlement   = "MaxDashboardCount"
	gatewayEntitlement     = "MaxGatewayCount"
)

func NewOrgListCommand(client internal.CloudClient) *cobra.Command {
	return NewCmd(list).
		WithExample("tykctl cloud org list --output<json/table>").
		WithLongDescription(orgListDesc).
		NoArgs(func(ctx context.Context, command cobra.Command) error {
			org, err := GetOrg(command.Context(), client)
			if err != nil {
				command.Println(err)
				return err
			}
			CreateOrgHeaderAndRows(org)
			return nil
		})
}

// GetOrg fetches all organization that belongs to a user.
func GetOrg(ctx context.Context, client internal.CloudClient) ([]cloud.Organisation, error) {
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
			organization.Name, organization.UID, getEntitlements(organization.Entitlements.Counters, teamEntitlement),
			getEntitlements(organization.Entitlements.Counters, environmentEntitlement), getEntitlements(organization.Entitlements.Counters, dashboardEntitlement),
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
