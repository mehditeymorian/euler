package component

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "component",
		Short: "generate graph for components",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {

}
