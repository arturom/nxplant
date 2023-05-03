package main

import (
	"fmt"
	"strings"
)

// type FieldSet map[string]string
type FieldSet map[string]interface{}

type RestSchema struct {
	Name   string   `json:"name"`
	Prefix string   `json:"@prefix"`
	Fields FieldSet `json:"fields"`
}

type SchemasResponse []RestSchema

type RestDocType struct {
	Parent  string   `json:"parent"`
	Facets  []string `json:"facets"`
	Schemas []string `json:"schemas"`
}

func (dt RestDocType) containsFacet(facet string) bool {
	for _, f := range dt.Facets {
		if f == facet {
			return true
		}
	}
	return false
}

func (dt RestDocType) isInvisible() bool {
	return dt.containsFacet("HiddenInNavigation")
}

type DocTypesMap map[string]RestDocType

type DocTypesResponse struct {
	DocTypes DocTypesMap `json:"doctypes"`
}

type SchemasMap = map[string]RestSchema

type RenderOptions struct {
	ExcludeOrphanSchemas bool
}

func RenderSchema(sb *strings.Builder, schema RestSchema) error {
	if _, err := sb.WriteString(fmt.Sprintf("abstract %s {\n", schema.Name)); err != nil {
		return err
	}
	for fieldName, fieldType := range schema.Fields {
		switch fieldType.(type) {
		case string:
			fieldType = fieldType
		default:
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

func RenderSchemas(sb *strings.Builder, schemas []RestSchema) error {
	for _, schema := range schemas {
		if err := RenderSchema(sb, schema); err != nil {
			return err
		}
	}
	return nil
}

func RenderDoctype(sb *strings.Builder, name string, doctype RestDocType) error {
	if _, err := sb.WriteString(fmt.Sprintf("class %s {\n}\n\n", name)); err != nil {
		return err
	}
	return nil
}

func RenderDocTypeRelations(sb *strings.Builder, docTypeName string, docType RestDocType) error {
	for _, schemaName := range docType.Schemas {
		if _, err := sb.WriteString(fmt.Sprintf("%s <|-- %s\n", docTypeName, schemaName)); err != nil {
			return err
		}
	}
	return nil
}

func RenderDocTypeParentRelation(sb *strings.Builder, name string, docType RestDocType) error {
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

func filterSchemas(schemas []RestSchema, usedSchemas map[string]bool) []RestSchema {
	result := make([]RestSchema, len(usedSchemas))
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

func RenderDocSchemas(sb *strings.Builder, schemas []RestSchema) error {
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

func RenderSchemasAndDocTypes(sb *strings.Builder, schemas []RestSchema, docTypes DocTypesMap, opts RenderOptions) error {
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
