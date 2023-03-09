package nxplant

import (
	"fmt"
	"strings"
)

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

func RenderSchema(sb *strings.Builder, schema Schema) error {
	if _, err := sb.WriteString(fmt.Sprintf("abstract %s {\n", schema.Name)); err != nil {
		return err
	}
	for fieldName, fieldType := range schema.Fields {
		if fieldType == "" {
			fieldType = "nested"
		}
		if _, err := sb.WriteString(fmt.Sprintf("   %s %s\n", fieldType, fieldName)); err != nil {
			return err
		}
	}
	if _, err := sb.WriteString("}\n\n"); err != nil {
		return err
	}
	return nil
}

func RenderSchemas(sb *strings.Builder, schemas []Schema) error {
	for _, schema := range schemas {
		if err := RenderSchema(sb, schema); err != nil {
			return err
		}
	}
	return nil
}

func RenderDoctype(sb *strings.Builder, name string, doctype DocType) error {
	if _, err := sb.WriteString(fmt.Sprintf("class %s {\n}\n\n", name)); err != nil {
		return err
	}
	return nil
}

func RenderDocTypeRelations(sb *strings.Builder, docTypeName string, docType DocType) error {
	for _, schemaName := range docType.Schemas {
		if _, err := sb.WriteString(fmt.Sprintf("%s <|-- %s\n", schemaName, docTypeName)); err != nil {
			return err
		}
	}
	return nil
}

func RenderDocTypeParentRelation(sb *strings.Builder, name string, docType DocType) error {
	if docType.Parent == "None!!!" {
		return nil
	}
	if _, err := sb.WriteString(fmt.Sprintf("%s <|-- %s\n", docType.Parent, name)); err != nil {
		return err
	}
	return nil
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

func RenderDocSchemas(sb *strings.Builder, schemas []Schema) error {
	if _, err := sb.WriteString("@startuml schemas\n\n"); err != nil {
		return err
	}
	RenderSchemas(sb, schemas)
	if _, err := sb.WriteString("@enduml\n"); err != nil {
		return err
	}
	return nil
}

func RenderDocTypes(sb *strings.Builder, docTypes DocTypesMap) error {
	if _, err := sb.WriteString("@startuml docTypes\n\n"); err != nil {
		return err
	}

	for docTypeName, docType := range docTypes {
		if err := RenderDoctype(sb, docTypeName, docType); err != nil {
			return err
		}
	}

	for docTypeName, docType := range docTypes {
		if err := RenderDocTypeParentRelation(sb, docTypeName, docType); err != nil {
			return err
		}
	}
	if _, err := sb.WriteString("\n@enduml\n"); err != nil {
		return err
	}
	return nil
}

func RenderSchemasAndDocTypes(sb *strings.Builder, schemas []Schema, docTypes DocTypesMap, opts RenderOptions) error {
	if _, err := sb.WriteString("@startuml schemasAndDocTypes\n\n"); err != nil {
		return err
	}
	if opts.ExcludeOrphanSchemas {
		schemas = filterSchemas(schemas, getUsedSchemas(docTypes))
	}
	if err := RenderSchemas(sb, schemas); err != nil {
		return nil
	}
	for docTypeName, docType := range docTypes {
		if err := RenderDoctype(sb, docTypeName, docType); err != nil {
			return err
		}
		if err := RenderDocTypeRelations(sb, docTypeName, docType); err != nil {
			return err
		}
	}
	for docTypeName, docType := range docTypes {
		if err := RenderDocTypeParentRelation(sb, docTypeName, docType); err != nil {
			return err
		}
	}

	if _, err := sb.WriteString("@enduml\n"); err != nil {
		return err
	}
	return nil
}
