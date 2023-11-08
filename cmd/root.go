package cmd

import (
	"github.com/mehditeymorian/euler/cmd/component"
	"github.com/mehditeymorian/euler/cmd/model"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "euler",
		Short: "euler graph generator",
		Long:  "euler graph generator",
	}

	rootCmd.AddCommand(
		model.Command(),
		component.Command(),
	)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
