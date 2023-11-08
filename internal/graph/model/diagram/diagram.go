package diagram

import (
	"context"
	"fmt"
	"github.com/mehditeymorian/euler/internal/graph/model/model"
	"io/ioutil"
	"path/filepath"
	"strings"

	"oss.terrastruct.com/d2"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2renderers/textmeasure"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
)

func CreateGraph(structs []model.Struct, withFields bool) string {
	components := new(strings.Builder)
	connections := new(strings.Builder)

	structsNames := make(map[string]bool)

	connectionExists := make(map[string]bool)

	// create structsNames
	for _, each := range structs {
		structsNames[each.Name] = true
	}

	for _, each := range structs {
		component := buildComponent(each, withFields)
		components.WriteString(component)

		for _, fieldType := range each.Fields {
			existKey := each.Name + "-" + fieldType
			exists, _ := connectionExists[existKey]
			if _, isStruct := structsNames[fieldType]; isStruct && !exists {
				connectionExists[existKey] = true
				connection := fmt.Sprintf("%s -> %s\n", each.Name, fieldType)
				connections.WriteString(connection)
			}
		}

	}

	components.WriteByte('\n')

	return components.String() + connections.String()
}

func Render(graph, address string) error {
	ruler, _ := textmeasure.NewRuler()
	diagram, err := d2.Compile(context.Background(), graph, &d2.CompileOptions{
		Layout:  d2elklayout.Layout,
		Ruler:   ruler,
		ThemeID: d2themescatalog.GrapeSoda.ID,
	})
	if err != nil {
		return err
	}

	out, _ := d2svg.Render(diagram)
	_ = ioutil.WriteFile(filepath.Join(address), out, 0600)
	return nil
}

func buildComponent(model model.Struct, withFields bool) string {
	if !withFields {
		return fmt.Sprintf("%s: %s\n", model.Name, model.Name)
	}

	component := new(strings.Builder)
	component.WriteString(fmt.Sprintf(`%s: {
	shape: class
`, model.Name))

	for fieldName, fieldType := range model.Fields {
		component.WriteString(fmt.Sprintf(`%s: "%s"
`, fieldName, fieldType))
	}

	component.WriteString(`}
`)

	return component.String()
}
