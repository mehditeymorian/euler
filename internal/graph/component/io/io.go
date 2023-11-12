package io

import (
	"fmt"
	"github.com/mehditeymorian/euler/internal/graph/component/model"
	"go/parser"
	"go/token"
	"golang.org/x/exp/slices"
	"golang.org/x/mod/modfile"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ScanComponents(path string, moduleName string, options model.Option) ([]model.Component, error) {
	components := make(map[string]map[string]bool)

	split := strings.Split(path, "/")
	pathPrefix := split[len(split)-1]

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			err := fmt.Errorf("failed to scan %s: %w", path, err)

			return err
		}

		if d.IsDir() || !strings.HasSuffix(d.Name(), ".go") {
			return nil
		}

		// Process Go source files
		sourceCode, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return nil
		}

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, sourceCode, parser.ImportsOnly)
		if err != nil {
			fmt.Println("Error parsing file:", err)
			return nil
		}

		packageName := filepath.Dir(path)
		split := strings.Split(packageName, "/")
		packageAbsoluteName := split[len(split)-1]

		packageName, _ = strings.CutPrefix(packageName, pathPrefix)

		key := buildKey(packageName, packageAbsoluteName)

		_, ok := components[key]
		if !ok {
			components[key] = make(map[string]bool)
		}

		for _, spec := range node.Imports {
			importPath, _ := strconv.Unquote(spec.Path.Value)
			internalDependency := strings.HasPrefix(importPath, moduleName)
			components[key][importPath] = internalDependency
		}

		return nil

	})
	if err != nil {
		return nil, fmt.Errorf("failed to scan files: %w", err)
	}

	var componentsArray []model.Component

	for componentName, importsMap := range components {
		packageName, packageAbsoluteName := breakKey(componentName)

		if slices.Contains(options.ExcludedComponents, componentName) {
			continue
		}

		component := model.Component{
			Name:         packageName,
			AbsoluteName: packageAbsoluteName,
		}

		var dependencies []model.Dependency

		for importName, internal := range importsMap {
			if slices.Contains(options.ExcludedDependencies, importName) {
				continue
			}

			absoluteName := importName
			if internal {
				absoluteName = strings.TrimPrefix(absoluteName, moduleName)
			}
			dependencies = append(dependencies, model.Dependency{
				Name:         importName,
				AbsoluteName: absoluteName,
				Internal:     internal,
			})
		}

		component.Dependencies = dependencies

		componentsArray = append(componentsArray, component)
	}

	return componentsArray, nil
}

func ModuleName(path string) (string, error) {
	// Read the content of the go.mod file
	modFileContent, err := ioutil.ReadFile(path + "/go.mod")
	if err != nil {
		return "", fmt.Errorf("error reading go.mod file: %w", err)
	}

	// Parse the go.mod file content
	modFile, err := modfile.Parse("go.mod", modFileContent, nil)
	if err != nil {
		return "", fmt.Errorf("error parsing go.mod file: %w", err)
	}

	// Extract the module name
	moduleName := modFile.Module.Mod.Path

	return moduleName, nil
}

func buildKey(packageName, absoluteName string) string {
	return packageName + ":" + absoluteName
}

func breakKey(key string) (string, string) {
	split := strings.Split(key, ":")
	packageName := split[0]
	packageAbsoluteName := split[1]

	return packageName, packageAbsoluteName
}
