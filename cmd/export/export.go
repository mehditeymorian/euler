package export

import (
	"encoding/json"
	"fmt"
	"github.com/mehditeymorian/euler/internal/graph/component/io"
	"github.com/mehditeymorian/euler/internal/graph/component/model"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "export components",
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
	excludedDependencies, _ := cmd.PersistentFlags().GetStringArray("exclude-dependencies")
	excludedComponents, _ := cmd.PersistentFlags().GetStringArray("exclude-components")
	cloned := false

	options := model.Option{
		ExcludedDependencies: excludedDependencies,
		ExcludedComponents:   excludedComponents,
	}

	if io.IsGitRepoURL(path) {
		fmt.Println("input project is a git repo. cloning...")
		tempPath, err := io.CloneRepository(path)
		if err != nil {
			panic(err)
		}
		path = tempPath

		fmt.Println("cloned repository successfully")

		defer os.RemoveAll(path) // Delete the temporary directory when done

		cloned = true
	}

	moduleName, err := io.ModuleName(path)
	if err != nil {
		panic(err)
	}

	components, err := io.ScanComponents(path, moduleName, cloned, options)
	if err != nil {
		panic(err)
	}

	marshal, err := json.Marshal(components)
	if err != nil {
		log.Printf("failed to marshal components: %v", err)
	}

	// write to file
	err = os.WriteFile("components.json", marshal, 0644)
	if err != nil {
		log.Printf("failed to write components to file: %v", err)

	}
}
