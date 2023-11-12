package component

import (
	"github.com/mehditeymorian/euler/internal/graph/component/diagram"
	"github.com/mehditeymorian/euler/internal/graph/component/io"
	"github.com/mehditeymorian/euler/internal/graph/component/model"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "component",
		Short: "generate graph for components",
		Run:   run,
	}

	cmd.PersistentFlags().StringP("path", "p", "./", "where to look for components")
	cmd.PersistentFlags().StringP("output", "o", "out.svg", "output file name")
	cmd.PersistentFlags().BoolP("render-external", "e", false, "rendering external dependencies too")
	cmd.PersistentFlags().StringArrayP("exclude-dependencies", "d", nil, "exclude these dependencies from output")
	cmd.PersistentFlags().StringArrayP("exclude-component", "c", nil, "exclude these components from output")

	return cmd
}

func run(cmd *cobra.Command, _ []string) {
	path, _ := cmd.PersistentFlags().GetString("path")
	outputFileName, _ := cmd.PersistentFlags().GetString("output")
	renderExternal, _ := cmd.PersistentFlags().GetBool("render-external")
	excludedDependencies, _ := cmd.PersistentFlags().GetStringArray("exclude-dependencies")
	excludedComponents, _ := cmd.PersistentFlags().GetStringArray("exclude-components")

	options := model.Option{
		ExcludedDependencies: excludedDependencies,
		ExcludedComponents:   excludedComponents,
	}

	moduleName, err := io.ModuleName(path)
	if err != nil {
		panic(err)
	}

	components, err := io.ScanComponents(path, moduleName, options)
	if err != nil {
		panic(err)
	}

	err = diagram.Render(components, moduleName, outputFileName, renderExternal)
	if err != nil {
		panic(err)
	}
}
