package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strings"

	lib "github.com/arturom/nxplant/lib"
	"github.com/arturom/nxplant/lib/osgi"
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

func readSchemas(filePath string) []lib.Schema {
	schemas := make([]lib.Schema, 0)
	readJsonFile(filePath, &schemas)
	return schemas
}

func readDocTypes(filePath string) lib.DocTypesResponse {
	docTypesResponse := lib.DocTypesResponse{}
	readJsonFile(filePath, &docTypesResponse)
	return docTypesResponse
}

func readComponent(filePath string) osgi.Component {
	component := osgi.Component{}
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
		if err := osgi.GenerateHierarchy(sb, component); err != nil {
			panic(err)
		}
	} else if *schemasFilePath != "" && *docTypesFilePath != "" {
		schemas := readSchemas(*schemasFilePath)
		docTypesResponse := readDocTypes(*docTypesFilePath)
		renderOptions := (lib.RenderOptions{
			ExcludeOrphanSchemas: true,
		})
		if err := lib.RenderSchemasAndDocTypes(sb, schemas, docTypesResponse.DocTypes, renderOptions); err != nil {
			panic(err)
		}
	} else if *schemasFilePath != "" {
		schemas := readSchemas(*schemasFilePath)
		if err := lib.RenderDocSchemas(sb, schemas); err != nil {
			panic(err)
		}
	} else if *docTypesFilePath != "" {
		docTypesResponse := readDocTypes(*docTypesFilePath)
		if err := lib.RenderDocTypes(sb, docTypesResponse.DocTypes); err != nil {
			panic(err)
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}
	fmt.Print(sb.String())
}
