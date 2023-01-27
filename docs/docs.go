package main

import (
	"bytes"
	"github.com/TykTechnologies/tykctl/cloudcmd"
	"github.com/TykTechnologies/tykctl/internal"
	"github.com/TykTechnologies/tykctl/sharedCmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
	"os"
)

func main() {
	rootCmd := sharedCmd.NewRootCmd()
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
	f, err := os.OpenFile("docs.md", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	out := new(bytes.Buffer)
	s := []*cobra.Command{cloudcmd.NewLoginCommand(factory),
		cloudcmd.NewInitCmd(factory),
		cloudcmd.NewOrgListCommand(factory),
		cloudcmd.NewCreateTeamCmd(factory),
		cloudcmd.NewFetchTeamCmd(factory),
	}
	for _, cmd := range s {
		err := doc.GenMarkdown(cmd, out)
		if err != nil {
			log.Fatal(err)
		}
		if _, err = f.Write(out.Bytes()); err != nil {
			panic(err)
		}
	}

}
