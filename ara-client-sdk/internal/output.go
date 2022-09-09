package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func showTable(headers []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetCenterSeparator("")
	table.SetAutoMergeCells(true)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderAlignment(3)

	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func showJson(body []byte) {
	var prettyJSON bytes.Buffer

	err := json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		log.Fatal("converting to json", err)
		///log.Debug("Raw output: ", string(body))
		return
	}

	fmt.Println(prettyJSON.String())
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
