package io

import (
	"errors"
	"fmt"
	"github.com/mehditeymorian/euler/internal/graph/model/model"
	"os"
	"regexp"
	"strings"
)

var ErrFailedToCompileRegex = errors.New("failed to compile regex")

const (
	StructPattern = "type\\s(?P<Name>\\w+)\\s+struct\\s+{(?P<Fields>[\\sa-zA-Z0-9*\\[\\]`\\.\\-'*\":]+)}"
)

func ExtractStructs(files []string) ([]model.Struct, error) {
	compile, err := regexp.Compile(StructPattern)
	if err != nil {
		return nil, ErrFailedToCompileRegex
	}

	var structs []model.Struct

	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("failed to read file at", file)
			continue
		}

		for _, eachStruct := range compile.FindAllSubmatch(bytes, -1) {
			name := string(eachStruct[1])
			fields := string(eachStruct[2])

			s := model.Struct{
				Name:   name,
				Fields: extractFields(fields),
			}

			structs = append(structs, s)
		}

	}

	return structs, nil
}

func extractFields(fieldsRaw string) map[string]string {
	trimmed := strings.TrimSpace(fieldsRaw)
	split := strings.Split(trimmed, "\n")

	fields := make(map[string]string)

	compile := regexp.MustCompile("\\s+")

	for _, each := range split {
		pair := compile.Split(strings.TrimSpace(each), -1)
		length := len(pair)
		var fieldName, fieldType string
		if length == 0 {
			continue
		} else if length == 1 {
			fieldName = pair[0]
			fieldType = pair[0]
		} else {
			fieldName = pair[0]
			fieldType = clearType(pair[1])
		}

		fields[fieldName] = fieldType
	}

	return fields
}

func clearType(typeRaw string) string {
	removeArray := strings.TrimPrefix(typeRaw, "[]")
	removePointer := strings.TrimPrefix(removeArray, "*")
	return removePointer
}
