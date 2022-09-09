package internal

import (
	"ara-client-sdk/swagger-gen"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

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

func PrintTeamInTable(teams []swagger.Team) {
	headers := []string{"Name", "ID", "Organization", "Frozen"}
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
