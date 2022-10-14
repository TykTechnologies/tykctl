package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"tykcli/swagger-gen"

	"github.com/olekukonko/tablewriter"
)

func ShowTable(headers []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetBorder(true)
	table.SetRowLine(true)
	//table.SetColumnSeparator("|")
	//table.SetCenterSeparator("")
	table.SetAutoMergeCells(false)
	//table.SetColumnSeparator("")
	//table.SetRowSeparator("")
	table.SetHeaderAlignment(3)
	table.SetAutoFormatHeaders(true)
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func ShowJson(body []byte) {
	var prettyJSON bytes.Buffer

	err := json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		log.Fatal("converting to json", err)
		///log.Debug("Raw output: ", string(body))
		return
	}

	fmt.Println(prettyJSON.String())
}

func PrintOrganizationTable(orgs []swagger.Organisation) {
	headers := []string{"Name", "Teams"}
	data := make([][]string, 0)
	for _, org := range orgs {
		orgData := []string{
			org.Name, strconv.Itoa(len(org.Teams)),
		}
		data = append(data, orgData)

	}
	ShowTable(headers, data)
}

func PrintDeploymentInTable(deployments []swagger.Deployment) {
	headers := []string{"ID", "Name", "State", "Control plane name", "Region", "Version", "Environment", "Team"}
	data := make([][]string, 0)
	for _, deployment := range deployments {
		deploymentData := []string{deployment.UID,
			deployment.Name, deployment.State, "", deployment.ZoneCode, deployment.BundleVersion, deployment.LoadoutName, deployment.TeamName,
		}
		data = append(data, deploymentData)

	}
	ShowTable(headers, data)

}

func PrintTeamInTable(teams []swagger.Team) {
	headers := []string{"Name", "ID", "Organisation", "Frozen"}
	data := make([][]string, 0)
	for _, team := range teams {
		teamData := []string{
			team.Name, team.UID, team.Organisation.Name, strconv.FormatBool(team.Frozen),
		}
		data = append(data, teamData)
	}

	ShowTable(headers, data)

}

func PrintLoadOutInTable(loadOut []swagger.Loadout) {
	headers := []string{"Name", "ID", "Team Name", "Blocked"}
	data := make([][]string, 0)
	for _, loadout := range loadOut {
		loadData := []string{
			loadout.Name, loadout.UID, loadout.TeamName, strconv.FormatBool(loadout.Blocked),
		}

		data = append(data, loadData)

	}
	ShowTable(headers, data)

}

func PrintOrganizationInTable(org []swagger.Organisation) {
	headers := []string{"Name", "ID", "Zone"}
	data := make([][]string, 0)
	for _, organisation := range org {
		orgData := []string{
			organisation.Name, organisation.UID, organisation.Zone,
		}

		data = append(data, orgData)

	}

	ShowTable(headers, data)

}

func tablePrintNoMerge(headers []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetCenterSeparator("")
	table.SetAutoMergeCells(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderAlignment(3)

	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
