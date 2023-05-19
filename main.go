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

func readSchemas(filePath string) SchemasResponse {
	schemas := make(SchemasResponse, 0)
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

func printDocumentation() {
	// components folder structure
	// components doc type inheritance
	// rest doc type inheritance
	// rest doc types and schemas
	// rest doc types with fields
	fmt.Fprintln(os.Stderr, "nxplant")
	fmt.Fprintln(os.Stderr, "  A Diagram generator for a Nuxeo project data model")
	fmt.Fprintln(os.Stderr)

	fmt.Fprintln(os.Stderr, "Examples of usage:")
	fmt.Fprintln(os.Stderr, "  nxplant --extensions extensions.xml")
	fmt.Fprintln(os.Stderr, "  nxplant --schemas schemas.json --types types.json")
	fmt.Fprintln(os.Stderr, "  nxplant --schemas schemas.json")
	fmt.Fprintln(os.Stderr, "  nxplant --types types.json")
	fmt.Fprintln(os.Stderr)
}

func writeDiagram(diagram PlantUMLDiagram, format string) {
	sb := &strings.Builder{}
	if format == "plantuml" {
		if err := diagram.writePlantuml(sb); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		if err := WriteD2(diagram, sb); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}
	fmt.Print(sb.String())
}

func main() {
	showHelp := flag.Bool("help", false, "Prints usage information")
	outputFormat := flag.String("format", "plantuml", "The diagram output format. Defaults to plantuml")
	extensionsFilePath := flag.String("extensions", "", "Path to XML file containing extensions")
	folderExtensionsFilePath := flag.String("folders", "", "Path to XML file containing extensions with a folder structure")
	schemasFilePath := flag.String("schemas", "", "Path to JSON file containing a list of schemas")
	docTypesFilePath := flag.String("types", "", "Path to JSON file containing the document types")
	flag.Parse()

	sb := &strings.Builder{}

	if *showHelp {
		printDocumentation()
		flag.Usage()
		os.Exit(0)
	}

	if *extensionsFilePath != "" {
		component := readComponent(*extensionsFilePath)
		diagram := GenerateDocumentHierarchyFromComponent(component)
		writeDiagram(*diagram, *outputFormat)
	} else if *folderExtensionsFilePath != "" {
		component := readComponent(*folderExtensionsFilePath)
		diagram := GenerateFolderStructureFromComponent(component)
		writeDiagram(*diagram, *outputFormat)
	} else if *schemasFilePath != "" && *docTypesFilePath != "" {
		docTypesResponse := readDocTypes(*docTypesFilePath)
		schemas := readSchemas(*schemasFilePath)
		diagram := GenerateTypesWithFields(docTypesResponse, schemas)
		writeDiagram(*diagram, *outputFormat)
	} else if *schemasFilePath != "" {
		schemas := readSchemas(*schemasFilePath)
		if err := RenderDocSchemas(sb, schemas); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	} else if *docTypesFilePath != "" {
		docTypesResponse := readDocTypes(*docTypesFilePath)
		diagram := GenerateTypesWithFacetsAndSchemas(docTypesResponse)
		writeDiagram(*diagram, *outputFormat)
	} else {
		printDocumentation()
		flag.Usage()
		os.Exit(1)
	}
	fmt.Print(sb.String())
}
