package main

import (
	"bytes"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	"github.com/TykTechnologies/tykctl/cloudcmd"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/sharedcmd"
)

func main() {
	rootCmd := sharedcmd.NewRootCmd()
	factory := internal.CloudFactory{
		Client: nil,
		Prompt: nil,
	}

	rootCmd.AddCommand(cloudcmd.NewCloudCommand(factory))
	rootCmd.AddCommand(cloudcmd.NewCtxCmd())

	err := doc.GenMarkdownTree(rootCmd, "./")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("docs.md", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	out := new(bytes.Buffer)
	s := []*cobra.Command{
		cloudcmd.NewLoginCommand(factory),
		cloudcmd.NewInitCmd(factory),
		cloudcmd.NewOrgListCommand(factory),
		cloudcmd.NewCreateTeamCmd(factory),
		cloudcmd.NewFetchTeamCmd(factory),
	}

	for _, cmd := range s {
		err := doc.GenMarkdown(cmd, out)
		if err != nil {
			panic(err)
		}

		if _, err = f.Write(out.Bytes()); err != nil {
			panic(err)
		}
	}
}
