package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
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

// PrintJson print an interface as json.
func PrintJson(body interface{}) error {
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