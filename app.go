package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	lib "github.com/arturom/nxplant/lib"
)

func readJSON(filePath string, obj interface{}) {
	text, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(text, &obj)
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

func main() {
	schemasFilePath := flag.String("schemas", "", "path to JSON file containing a list of schemas")
	docTypesFilePath := flag.String("types", "", "Path to JSON file containing the document types")
	outputFilePath := flag.String("out", "", "Path to file output")
	flag.Parse()

	if *outputFilePath == "" {
		flag.CommandLine.Usage()
		panic("Missing parameters")
	}

	renderOptions := (lib.RenderOptions{
		ExcludeOrphanSchemas: true,
	})

	if *schemasFilePath != "" && *docTypesFilePath != "" {
		fmt.Println("one")
		schemas := readSchemas(*schemasFilePath)
		docTypesResponse := readDocTypes(*docTypesFilePath)
		result := lib.RenderSchemasAndDocTypes(schemas, docTypesResponse.DocTypes, renderOptions)
		ioutil.WriteFile(*outputFilePath, []byte(result), 0644)
		return
	}

	if *schemasFilePath != "" {
		fmt.Println("two")
		schemas := readSchemas(*schemasFilePath)
		result := lib.RenderDocSchemas(schemas)
		ioutil.WriteFile(*outputFilePath, []byte(result), 0644)
		return
	}

	if *docTypesFilePath != "" {
		fmt.Println("three")
		docTypesResponse := readDocTypes(*docTypesFilePath)
		result := lib.RenderDocTypes(docTypesResponse.DocTypes)
		ioutil.WriteFile(*outputFilePath, []byte(result), 0644)
		return
	}

	f, err := os.Create(*outputFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString("hello there\n")
	f.WriteString("general kenobi\n")

}
