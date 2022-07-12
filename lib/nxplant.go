package nxplant

import "fmt"

func parse() {
	fmt.Println("")
}

const NL = "\n"

type FieldSet map[string]string

type Schema struct {
	Name   string   `json:"name"`
	Prefix string   `json:"@prefix"`
	Fields FieldSet `json:"fields"`
}

type DocType struct {
	Parent  string   `json:"parent"`
	Facets  []string `json:"facets"`
	Schemas []string `json:"schemas"`
}

type DocTypesMap = map[string]DocType

type DocTypesResponse struct {
	DocTypes DocTypesMap `json:"doctypes"`
}

type SchemasMap = map[string]Schema

type RenderOptions struct {
	ExcludeOrphanSchemas bool
}

func RenderSchema(schema Schema) string {
	result := fmt.Sprintf("abstract %s {\n", schema.Name)
	for fieldName, fieldType := range schema.Fields {
		if fieldType == "" {
			fieldType = "nested"
		}
		result += fmt.Sprintf("   %s %s\n", fieldType, fieldName)
	}
	result += "}\n\n"
	return result
}

func RenderSchemas(schemas []Schema) string {
	result := ""
	for _, schema := range schemas {
		result += RenderSchema(schema)
	}
	return result
}

func RenderDoctype(name string, doctype DocType) string {
	result := fmt.Sprintf("class %s {\n", name)
	result += "}\n"
	return result
}

func RenderDocTypeRelations(docTypeName string, docType DocType) string {
	result := ""
	for _, schemaName := range docType.Schemas {
		result += fmt.Sprintf("%s <|-- %s\n", schemaName, docTypeName)
	}
	return result
}

func RenderDocTypeParentRelation(name string, docType DocType) string {
	if docType.Parent == "None!!!" {
		return ""
	}
	return fmt.Sprintf("%s <|-- %s\n", docType.Parent, name)
}

func getUsedSchemas(docTypes DocTypesMap) map[string]bool {
	result := make(map[string]bool)
	for _, docType := range docTypes {
		for _, schema := range docType.Schemas {
			result[schema] = true
		}
	}
	return result
}

func filterSchemas(schemas []Schema, usedSchemas map[string]bool) []Schema {
	result := make([]Schema, len(usedSchemas))
	n := 0
	for _, schema := range schemas {
		exists, _ := usedSchemas[schema.Name]
		if exists {
			result[n] = schema
			n++
		}
	}
	return result
}

func RenderDocSchemas(schemas []Schema) string {
	result := "@startuml schemas\n\n"
	result += RenderSchemas(schemas)
	result += "@enduml\n"

	return result
}

func RenderDocTypes(docTypes DocTypesMap) string {
	result := "@startuml docTypes\n\n"

	for docTypeName, docType := range docTypes {
		result += RenderDoctype(docTypeName, docType)
		result += NL
	}

	for docTypeName, docType := range docTypes {
		result += RenderDocTypeParentRelation(docTypeName, docType)
	}
	result += NL
	result += "@enduml\n"
	return result
}

func RenderSchemasAndDocTypes(schemas []Schema, docTypes DocTypesMap, opts RenderOptions) string {
	result := "@startuml schemasAndDocTypes\n\n"
	if opts.ExcludeOrphanSchemas {
		schemas = filterSchemas(schemas, getUsedSchemas(docTypes))
	}
	result += RenderSchemas(schemas)
	result += NL
	for docTypeName, docType := range docTypes {
		result += RenderDoctype(docTypeName, docType)
		result += NL
		result += RenderDocTypeRelations(docTypeName, docType)
		result += NL
	}
	result += NL

	for docTypeName, docType := range docTypes {
		result += RenderDocTypeParentRelation(docTypeName, docType)
	}
	result += NL

	result += "@enduml\n"
	return result
}
