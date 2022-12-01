package io

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var ErrFailedToCompileRegex = errors.New("failed to compile regex")

const (
	StructPattern = "type\\s(?P<Name>\\w+)\\sstruct\\s{(?P<Fields>(\\s|\\w)+)}"
)

type Struct struct {
	Name   string
	Fields map[string]string
}

func ExtractStructs(files []string) ([]Struct, error) {
	compile, err := regexp.Compile(StructPattern)
	if err != nil {
		return nil, ErrFailedToCompileRegex
	}

	var structs []Struct

	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("failed to read file at", file)
			continue
		}

		for _, eachStruct := range compile.FindAllSubmatch(bytes, -1) {
			name := string(eachStruct[1])
			fields := string(eachStruct[2])

			s := Struct{
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
		fieldName := pair[0]
		fieldType := pair[1]

		fields[fieldName] = fieldType
	}

	return fields
}
