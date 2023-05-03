package main

import "fmt"

func GenerateFolderStructureFromComponent(c Component) *Diagram {
	diagram := &Diagram{
		Name:      "Folder Structure",
		Classes:   make([]Class, 0),
		Relations: make([]Relation, 0),
	}
	existingRels := make(map[string]bool)

	for _, ext := range c.Extensions {
		if len(ext.DocTypes) == 0 {
			// exclude components that are not schema related
			continue
		}
		for _, docType := range ext.DocTypes {
			/*
				class := Class{
					Name:   docType.Name,
					Fields: make([]Field, 0),
				}
				diagram.Classes = append(diagram.Classes, class)
			*/

			/*
				if docType.Extends != "" {
					rel := Relation{
						From:   docType.Extends,
						To:     docType.Name,
						Label:  "extends",
						Symbol: "..",
					}
					diagram.Relations = append(diagram.Relations, rel)
				}
			*/

			for _, subType := range docType.Subtypes.Types {
				searchKey := fmt.Sprintf("%s:%s", docType.Name, subType)
				if _, found := existingRels[searchKey]; found {
					continue
				} else {
					existingRels[searchKey] = true
				}
				rel := Relation{
					From: docType.Name,
					To:   subType,
				}
				diagram.Relations = append(diagram.Relations, rel)
			}

		}
	}
	return diagram
}

func GenerateDocumentHierarchyFromComponent(c Component) *Diagram {
	diagram := &Diagram{
		Name:      "Custom Document Types",
		Classes:   make([]Class, 0),
		Relations: make([]Relation, 0),
	}

	for _, ext := range c.Extensions {
		if len(ext.DocTypes) == 0 {
			// exclude components that are not schema related
			continue
		}
		for _, docType := range ext.DocTypes {
			class := Class{
				Type:    "class",
				Name:    docType.Name,
				content: make([]Writable, 0),
			}
			diagram.Classes = append(diagram.Classes, class)

			if docType.Extends != "" {
				rel := Relation{
					From:  docType.Extends,
					To:    docType.Name,
					Label: "extends",
				}
				diagram.Relations = append(diagram.Relations, rel)
			}

			for _, subType := range docType.Subtypes.Types {
				field := Field{
					Name: subType,
				}
				class.content = append(class.content, field)
			}

		}
	}

	return diagram
}

func GenerateTypesWithFacetsAndSchemas(docTypes DocTypesResponse) *Diagram {
	diagram := &Diagram{
		Name:      "Document Types",
		Classes:   make([]Class, 0),
		Relations: make([]Relation, 0),
	}

	for name, docType := range docTypes.DocTypes {
		if docType.isInvisible() {
			continue
		}
		class := Class{
			Type:    "class",
			Name:    name,
			content: make([]Writable, 0),
		}

		class.content = append(class.content, Separation{val: "facets"})
		for _, facetName := range docType.Facets {
			field := Field{
				Name: facetName,
			}
			class.content = append(class.content, field)
		}

		class.content = append(class.content, Separation{val: "schemas"})
		for _, schema := range docType.Schemas {
			field := Field{
				Char: "#",
				Name: schema,
			}
			class.content = append(class.content, field)
		}
		diagram.Classes = append(diagram.Classes, class)

		if docType.Parent != "None!!!" {
			rel := Relation{
				From: docType.Parent,
				To:   name,
			}

			diagram.Relations = append(diagram.Relations, rel)
		}
	}

	return diagram
}

func GenerateTypesWithFields(docTypes DocTypesResponse, schemas SchemasResponse) *Diagram {
	diagram := &Diagram{
		Name:      "Document Types and Fields",
		Classes:   make([]Class, 0),
		Relations: make([]Relation, 0),
	}

	schemasMap := make(map[string]RestSchema)
	for _, schema := range schemas {
		schemasMap[schema.Name] = schema
	}

	for name, docType := range docTypes.DocTypes {
		if docType.isInvisible() {
			continue
		}
		class := Class{
			Type:    "class",
			Name:    name,
			content: make([]Writable, 0),
		}

		class.content = append(class.content, Separation{val: "facets"})
		for _, facetName := range docType.Facets {
			field := Field{
				Name: facetName,
			}
			class.content = append(class.content, field)
		}

		class.content = append(class.content, Separation{val: "schemas"})
		for _, schema := range docType.Schemas {
			field := Field{
				Char: "#",
				Name: schema,
			}
			class.content = append(class.content, field)
		}

		for _, schemaName := range docType.Schemas {
			schema, _ := schemasMap[schemaName]
			class.content = append(class.content, Separation{val: schemaName})
			for fieldName, fieldType := range schema.Fields {
				switch fieldType.(type) {
				case string:
					fieldType = fieldType
				default:
					fieldType = "nested"
				}
				field := Field{
					Type: fieldType.(string),
					Name: fieldName,
				}
				class.content = append(class.content, field)
			}
		}

		diagram.Classes = append(diagram.Classes, class)

		if docType.Parent != "None!!!" {
			rel := Relation{
				From: docType.Parent,
				To:   name,
			}

			diagram.Relations = append(diagram.Relations, rel)
		}
	}

	return diagram
}
