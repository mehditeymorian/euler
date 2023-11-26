package model

import (
	"github.com/mehditeymorian/euler/internal/graph/model/diagram"
	io2 "github.com/mehditeymorian/euler/internal/graph/model/io"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "struct",
		Short: "struct graph generator",
		Long:  "generate graph for struct models",
		Run:   run,
	}

	cmd.PersistentFlags().StringP("path", "p", "./", "directory of model")
	cmd.PersistentFlags().StringSliceP("exclude", "e", nil, "exclude files and directories in scanning")
	cmd.PersistentFlags().BoolP("fields", "f", false, "include structs fields in diagram")

	return cmd
}

func run(cmd *cobra.Command, _ []string) {
	path, _ := cmd.PersistentFlags().GetString("path")
	excludes, _ := cmd.PersistentFlags().GetStringSlice("exclude")
	withFields, _ := cmd.PersistentFlags().GetBool("fields")

	files, err := io2.Files(path, excludes)
	if err != nil {
		panic("failed to read files in the given directory")
	}

	structs, err := io2.ExtractStructs(files)
	if err != nil {
		panic(err)
	}

	graph := diagram.CreateGraph(structs, withFields)

	err = diagram.Render(graph, "out.svg")
	if err != nil {
		panic(err)
	}

}
