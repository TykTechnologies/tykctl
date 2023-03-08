package cloudcmd

import (
	"context"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
)

const zonesDesc = `
This will fetch all the zones and the type of deployment they support.

You can run this command to find out the regions you can deploy your home or edge deployments.
`

func NewZonesCmd(factory internal.CloudFactory) *cobra.Command {
	return internal.NewCmd(zones).
		AddPreRunFuncs(NewCloudRbac(TeamMember, factory.Config).CloudRbac).
		WithDescription("fetch all the supported zones and what deployment they support.").
		WithExample("tykctl cloud zones").
		WithLongDescription(zonesDesc).
		WithFlagAdder(false, addOutPutFlags).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			err := FetchZonesAndPrint(ctx, factory.Client, cmd.Flags())
			if err != nil {
				cmd.PrintErrln(err)
				return err
			}
			return nil
		})
}

func FetchZonesAndPrint(ctx context.Context, client internal.CloudClient, f *pflag.FlagSet) error {
	format, err := f.GetString(outPut)
	if err != nil {
		return err
	}
	if format != table && format != jsonFormat {
		return ErrorOutPutFormat
	}
	deploymentZones, _, err := client.GetDeploymentZones(ctx)
	if err != nil {
		return err
	}
	if format == table {
		internal.Printable(ZonesTable(deploymentZones.Payload))
		return nil
	}
	return internal.PrintJSON(deploymentZones)
}

func ZonesTable(response internal.Payload) ([]string, [][]string) {
	header := []string{"Name", "Support Home", "Support Gateway"}
	rows := make([][]string, 0)
	for s, supported := range response.Tags {
		row := []string{
			s,
			strconv.FormatBool(util.Contains(supported, "Home")),
			strconv.FormatBool(util.Contains(supported, "Gateway")),
		}
		rows = append(rows, row)
	}
	return header, rows
}
