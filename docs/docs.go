package main

import (
	"github.com/TykTechnologies/tykctl/cloudcmd"
	"github.com/spf13/cobra/doc"
	"log"
)

func main() {
	tykctl := cloudcmd.NewCloudCommand(nil)
	err := doc.GenMarkdownTree(tykctl, "./")
	if err != nil {
		log.Fatal(err)
	}
}
