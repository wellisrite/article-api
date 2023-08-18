package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd := cobra.Command{
		Use:   "app",
		Short: "Article API",
		Run: func(*cobra.Command, []string) {
			serv()
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "run-server",
		Short: "Run Article API Server",
		Run: func(*cobra.Command, []string) {
			serv()
		},
	})

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
