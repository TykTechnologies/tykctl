package main

import (
	"bytes"
	"github.com/TykTechnologies/tykctl/cloudcmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("docs.md", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	out := new(bytes.Buffer)
	s := []*cobra.Command{cloudcmd.NewLoginCommand()}
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
