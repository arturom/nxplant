package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/arturom/nxplant/diagrams"
	components "github.com/arturom/nxplant/input/components/doctypes"
	rest "github.com/arturom/nxplant/input/restapi/doctypes"
	"github.com/arturom/nxplant/output/d2"
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

func readSchemas(filePath string) rest.SchemasResponse {
	schemas := make(rest.SchemasResponse, 0)
	readJsonFile(filePath, &schemas)
	return schemas
}

func readDocTypes(filePath string) rest.DocTypesResponse {
	docTypesResponse := rest.DocTypesResponse{}
	readJsonFile(filePath, &docTypesResponse)
	return docTypesResponse
}

func readComponent(filePath string) components.Component {
	component := components.Component{}
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

func writeDiagram(diagram diagrams.PlantUMLDiagram, format string) {
	sb := &strings.Builder{}
	if format == "plantuml" {
		if err := diagram.WritePlantuml(sb); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		if err := d2.WriteD2(diagram, sb); err != nil {
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
		diagram := components.GenerateDocumentHierarchyFromComponent(component)
		writeDiagram(*diagram, *outputFormat)
	} else if *folderExtensionsFilePath != "" {
		component := readComponent(*folderExtensionsFilePath)
		diagram := components.GenerateFolderStructureFromComponent(component)
		writeDiagram(*diagram, *outputFormat)
	} else if *schemasFilePath != "" && *docTypesFilePath != "" {
		docTypesResponse := readDocTypes(*docTypesFilePath)
		schemas := readSchemas(*schemasFilePath)
		diagram := rest.GenerateTypesWithFields(docTypesResponse, schemas)
		writeDiagram(*diagram, *outputFormat)
	} else if *schemasFilePath != "" {
		schemas := readSchemas(*schemasFilePath)
		if err := rest.RenderDocSchemas(sb, schemas); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	} else if *docTypesFilePath != "" {
		docTypesResponse := readDocTypes(*docTypesFilePath)
		diagram := rest.GenerateTypesWithFacetsAndSchemas(docTypesResponse)
		writeDiagram(*diagram, *outputFormat)
	} else {
		printDocumentation()
		flag.Usage()
		os.Exit(1)
	}
	fmt.Print(sb.String())
}
