package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/tidwall/gjson"
)

// Printable will print the data as a table on the terminal.
func Printable(headers []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.AppendBulk(data)
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoMergeCells(false)
	table.SetHeaderAlignment(3)
	table.SetAutoFormatHeaders(true)
	table.Render()
}

func BorderLessTable(headers []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(headers)
	table.AppendBulk(data)
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetAutoMergeCells(false)
	table.SetHeaderAlignment(3)
	table.SetAutoFormatHeaders(false)
	table.SetCenterSeparator("")
	table.SetRowSeparator("")
	table.SetColumnSeparator("")
	table.Render()
}

// PrintJSON print an interface as json.
func PrintJSON(body interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	var prettyJSON bytes.Buffer

	err = json.Indent(&prettyJSON, b, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(prettyJSON.String())

	return nil
}

func HandleGets(body interface{}, keys []string) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	hdrs := []string{"Key", "Data"}
	rows := make([][]string, 0)

	results := gjson.GetManyBytes(b, keys...)
	for i, result := range results {
		row := []string{keys[i], result.Raw}
		rows = append(rows, row)
	}

	BorderLessTable(hdrs, rows)

	return nil
}
