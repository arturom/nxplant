package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strings"
)

func readJsonFile(filePath string, obj interface{}) {
	text, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(text, &obj)
	if err != nil {
		panic(err)
	}
}

func readXmlFile(filePath string, obj interface{}) {
	text, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(text, &obj)
	if err != nil {
		panic(err)
	}
}

func readSchemas(filePath string) []RestSchema {
	schemas := make([]RestSchema, 0)
	readJsonFile(filePath, &schemas)
	return schemas
}

func readDocTypes(filePath string) DocTypesResponse {
	docTypesResponse := DocTypesResponse{}
	readJsonFile(filePath, &docTypesResponse)
	return docTypesResponse
}

func readComponent(filePath string) Component {
	component := Component{}
	readXmlFile(filePath, &component)
	return component
}

func main() {
	extensionsFilePath := flag.String("extensions", "", "path to XML file containing extensions")
	schemasFilePath := flag.String("schemas", "", "path to JSON file containing a list of schemas")
	docTypesFilePath := flag.String("types", "", "Path to JSON file containing the document types")
	flag.Parse()

	sb := &strings.Builder{}

	if *extensionsFilePath != "" {
		component := readComponent(*extensionsFilePath)
		if err := GenerateHierarchy(sb, component); err != nil {
			panic(err)
		}
	} else if *schemasFilePath != "" && *docTypesFilePath != "" {
		schemas := readSchemas(*schemasFilePath)
		docTypesResponse := readDocTypes(*docTypesFilePath)
		renderOptions := (RenderOptions{
			ExcludeOrphanSchemas: true,
		})
		if err := RenderSchemasAndDocTypes(sb, schemas, docTypesResponse.DocTypes, renderOptions); err != nil {
			panic(err)
		}
	} else if *schemasFilePath != "" {
		schemas := readSchemas(*schemasFilePath)
		if err := RenderDocSchemas(sb, schemas); err != nil {
			panic(err)
		}
	} else if *docTypesFilePath != "" {
		docTypesResponse := readDocTypes(*docTypesFilePath)
		if err := RenderDocTypes(sb, docTypesResponse.DocTypes); err != nil {
			panic(err)
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}
	fmt.Print(sb.String())
}
