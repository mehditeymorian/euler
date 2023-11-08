package component

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "component",
		Short: "generate graph for components",
		Run:   run,
	}

	cmd.PersistentFlags().StringP("path", "p", "./", "where to look for components")

	return cmd
}

func run(cmd *cobra.Command, _ []string) {
	path, _ := cmd.PersistentFlags().GetString("path")

	fmt.Println(path)
}
