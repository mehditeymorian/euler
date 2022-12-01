package cmd

import (
	"os"

	"github.com/mehditeymorian/euler/internal/diagram"
	"github.com/mehditeymorian/euler/internal/io"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "euler",
		Short: "euler graph generator",
		Long:  "euler graph generator for golang struct model",
		Run:   run,
	}

	rootCmd.PersistentFlags().StringP("path", "p", "./", "directory of model")
	rootCmd.PersistentFlags().StringSliceP("exclude", "e", nil, "exclude files and directories in scanning")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, _ []string) {
	path, _ := cmd.PersistentFlags().GetString("path")
	excludes, _ := cmd.PersistentFlags().GetStringSlice("exclude")

	files, err := io.Files(path, excludes)
	if err != nil {
		panic("failed to read files in the given directory")
	}

	structs, err := io.ExtractStructs(files)
	if err != nil {
		panic(err)
	}

	graph := diagram.CreateGraph(structs)

	err = diagram.Render(graph, "out.svg")
	if err != nil {
		panic(err)
	}

}
