package diagram

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/mehditeymorian/euler/internal/model"
	"oss.terrastruct.com/d2"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2renderers/textmeasure"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
)

func CreateGraph(structs []model.Struct) string {
	components := new(strings.Builder)
	connections := new(strings.Builder)

	structsNames := make(map[string]bool)

	connectionExists := make(map[string]bool)

	// create structsNames
	for _, each := range structs {
		structsNames[each.Name] = true
	}

	for _, each := range structs {
		component := fmt.Sprintf("%s: %s\n", each.Name, each.Name)
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
		Layout:  d2dagrelayout.Layout,
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
