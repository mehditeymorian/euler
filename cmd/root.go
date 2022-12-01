package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "euler",
		Short: "euler graph generator",
		Long:  "euler graph generator for golang struct models",
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


