package diagram

import (
	"context"
	"fmt"
	"github.com/mehditeymorian/euler/internal/graph/component/model"
	"io/ioutil"
	"oss.terrastruct.com/d2"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2renderers/textmeasure"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
	"path/filepath"
	"strings"
)

func Render(components []model.Component, moduleName, fileName string, renderExternalDependencies bool) error {
	d2Data := generateD2(components, moduleName, renderExternalDependencies)

	ruler, err := textmeasure.NewRuler()
	if err != nil {
		return fmt.Errorf("failed to create new ruler: %w", err)
	}
	diagram, err := d2.Compile(context.Background(), d2Data, &d2.CompileOptions{
		Layout:  d2elklayout.Layout,
		Ruler:   ruler,
		ThemeID: d2themescatalog.GrapeSoda.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to compile data: %w", err)
	}

	out, err := d2svg.Render(diagram)
	if err != nil {
		return fmt.Errorf("failed to render diagram: %w", err)
	}

	err = ioutil.WriteFile(filepath.Join(fileName), out, 0600)
	if err != nil {
		return fmt.Errorf("failed to write diagram to file: %w", err)
	}

	return nil
}

func generateD2(components []model.Component, moduleName string, renderExternalDependencies bool) string {
	diagram := new(strings.Builder)

	write := func(format string, args ...any) {
		diagram.WriteString(fmt.Sprintf(format, args...))
	}

	for _, component := range components {
		for _, dependency := range component.Dependencies {
			if dependency.Internal {
				componentName := cutPrefixes(component.Name, moduleName, "/")
				componentName = strings.ReplaceAll(componentName, "/", ".")

				dependencyName := cutPrefixes(dependency.Name, moduleName, "/")
				dependencyName = strings.ReplaceAll(dependencyName, "/", ".")

				if componentName == "" || dependencyName == "" || componentName == dependencyName {
					continue
				}

				write("%s -> %s\n", dependencyName, componentName)
			} else if renderExternalDependencies {
				dependencyName := strings.ReplaceAll(dependency.Name, ".", "/")
				componentName := cutPrefixes(component.Name, moduleName, "/")
				componentName = strings.ReplaceAll(componentName, "/", ".")

				if componentName == "" || dependencyName == "" || componentName == dependencyName {
					continue
				}

				write("%s -> %s\n", dependencyName, componentName)
			}
		}
	}

	return diagram.String()
}

func cutPrefixes(input string, prefixes ...string) string {
	result := input
	for _, prefix := range prefixes {
		after, _ := strings.CutPrefix(result, prefix)
		result = after
	}

	return result
}
