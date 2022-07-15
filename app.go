package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"os"

	lib "github.com/arturom/nxplant/lib"
	"github.com/arturom/nxplant/lib/osgi"
)

func readJSON(filePath string, obj interface{}) {
	text, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(text, &obj)
	if err != nil {
		panic(err)
	}
}

func readXML(filePath string, obj interface{}) {
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
	readJSON(filePath, &schemas)
	return schemas
}

func readDocTypes(filePath string) lib.DocTypesResponse {
	docTypesResponse := lib.DocTypesResponse{}
	readJSON(filePath, &docTypesResponse)
	return docTypesResponse
}

func readComponent(filePath string) osgi.Component {
	component := osgi.Component{}
	readXML(filePath, &component)
	return component
}

func main() {
	schemasFilePath := flag.String("schemas", "", "path to JSON file containing a list of schemas")
	docTypesFilePath := flag.String("types", "", "Path to JSON file containing the document types")
	outputFilePath := flag.String("out", "", "Path to file output")
	flag.Parse()

	if *outputFilePath == "" {
		// flag.CommandLine.Usage()
		// panic("Missing parameters")
	}

	renderOptions := (lib.RenderOptions{
		ExcludeOrphanSchemas: true,
	})

	if *schemasFilePath != "" && *docTypesFilePath != "" {
		schemas := readSchemas(*schemasFilePath)
		docTypesResponse := readDocTypes(*docTypesFilePath)
		result := lib.RenderSchemasAndDocTypes(schemas, docTypesResponse.DocTypes, renderOptions)
		ioutil.WriteFile(*outputFilePath, []byte(result), 0644)
		return
	}

	if *schemasFilePath != "" {
		schemas := readSchemas(*schemasFilePath)
		result := lib.RenderDocSchemas(schemas)
		ioutil.WriteFile(*outputFilePath, []byte(result), 0644)
		return
	}

	if *docTypesFilePath != "" {
		docTypesResponse := readDocTypes(*docTypesFilePath)
		result := lib.RenderDocTypes(docTypesResponse.DocTypes)
		ioutil.WriteFile(*outputFilePath, []byte(result), 0644)
		return
	}

	component := readComponent("/Users/arturomejia/projects/nxplant/extensions.xml")
	osgi.GenerateHierarchy(component)
}
