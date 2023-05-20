package doctypes

import (
	"github.com/arturom/nxplant/diagrams"
)

func GenerateTypesWithFacetsAndSchemas(docTypes DocTypesResponse) *diagrams.Diagram {
	diagram := &diagrams.Diagram{
		Name:      "Document Types",
		Classes:   make([]diagrams.Class, 0),
		Relations: make([]diagrams.Relation, 0),
	}

	for name, docType := range docTypes.DocTypes {
		if docType.IsInvisible() {
			continue
		}
		class := diagrams.Class{
			Type:    "class",
			Name:    name,
			Content: make([]diagrams.Writable, 0),
		}

		class.Content = append(class.Content, diagrams.Separation{Val: "facets"})
		for _, facetName := range docType.Facets {
			field := diagrams.Field{
				Name: facetName,
			}
			class.Content = append(class.Content, field)
		}

		class.Content = append(class.Content, diagrams.Separation{Val: "schemas"})
		for _, schema := range docType.Schemas {
			field := diagrams.Field{
				Char: "#",
				Name: schema,
			}
			class.Content = append(class.Content, field)
		}
		diagram.Classes = append(diagram.Classes, class)

		if docType.HasParent() {
			rel := diagrams.Relation{
				From: docType.Parent,
				To:   name,
			}

			diagram.Relations = append(diagram.Relations, rel)
		}
	}

	return diagram
}

func GenerateTypesWithFields(docTypes DocTypesResponse, schemas SchemasResponse) *diagrams.Diagram {
	diagram := &diagrams.Diagram{
		Name:      "Document Types and Fields",
		Classes:   make([]diagrams.Class, 0),
		Relations: make([]diagrams.Relation, 0),
	}

	schemasMap := make(map[string]RestSchema)
	for _, schema := range schemas {
		schemasMap[schema.Name] = schema
	}

	for name, docType := range docTypes.DocTypes {
		if docType.IsInvisible() {
			continue
		}
		class := diagrams.Class{
			Type:    "class",
			Name:    name,
			Content: make([]diagrams.Writable, 0),
		}

		class.Content = append(class.Content, diagrams.Separation{Val: "facets"})
		for _, facetName := range docType.Facets {
			field := diagrams.Field{
				Name: facetName,
			}
			class.Content = append(class.Content, field)
		}

		class.Content = append(class.Content, diagrams.Separation{Val: "schemas"})
		for _, schema := range docType.Schemas {
			field := diagrams.Field{
				Char: "#",
				Name: schema,
			}
			class.Content = append(class.Content, field)
		}

		for _, schemaName := range docType.Schemas {
			schema, _ := schemasMap[schemaName]
			class.Content = append(class.Content, diagrams.Separation{Val: schemaName})
			for fieldName, fieldType := range schema.Fields {
				switch fieldType.(type) {
				case string:
					fieldType = fieldType
				default:
					fieldType = "nested"
				}
				field := diagrams.Field{
					Type: fieldType.(string),
					Name: fieldName,
				}
				class.Content = append(class.Content, field)
			}
		}

		diagram.Classes = append(diagram.Classes, class)

		if docType.HasParent() {
			rel := diagrams.Relation{
				From: docType.Parent,
				To:   name,
			}

			diagram.Relations = append(diagram.Relations, rel)
		}
	}

	return diagram
}
