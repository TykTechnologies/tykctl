package cmd

import (
	"context"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strconv"
)

func NewZonesCmd(client internal.CloudClient) *cobra.Command {
	return NewCmd(zones).
		WithFlagAdder(false, addOutPutFlags).
		NoArgs(func(ctx context.Context, cmd cobra.Command) error {
			err := FetchZonesAndPrint(ctx, client, cmd.Flags())
			if err != nil {
				cmd.Println(err)
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
	return internal.PrintJson(deploymentZones)
}

func ZonesTable(response internal.Payload) ([]string, [][]string) {
	header := []string{"Name", "Support Home", "Support Gateway"}
	rows := make([][]string, 0)
	for s, supported := range response.Tags {
		row := []string{s,
			strconv.FormatBool(util.Contains(supported, "Home")),
			strconv.FormatBool(util.Contains(supported, "Gateway")),
		}
		rows = append(rows, row)
	}
	return header, rows
}
