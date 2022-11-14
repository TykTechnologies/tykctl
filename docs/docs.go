package main

import (
	"github.com/TykTechnologies/tykctl/cmd"
	"github.com/spf13/cobra/doc"
	"log"
)

func main() {
	tykctl := cmd.NewRootCmd(nil)
	err := doc.GenMarkdownTree(tykctl, "./")
	if err != nil {
		log.Fatal(err)
	}
}
