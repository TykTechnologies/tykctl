package main

import (
	"bytes"
	"github.com/TykTechnologies/tykctl/cloudcmd"
	"github.com/TykTechnologies/tykctl/sharedCmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
	"os"
)

func main() {
	rootCmd := sharedCmd.NewRootCmd()
	rootCmd.AddCommand(cloudcmd.NewCloudCommand(nil))
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
	s := []*cobra.Command{cloudcmd.NewLoginCommand(nil),
		cloudcmd.NewInitCmd(nil),
		cloudcmd.NewOrgListCommand(nil),
		cloudcmd.NewCreateTeamCmd(nil),
		cloudcmd.NewFetchTeamCmd(nil),
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
