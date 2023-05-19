package doctypes

import (
	"fmt"

	"github.com/arturom/nxplant/diagrams"
)

func GenerateFolderStructure(c Component) *diagrams.PlantUMLDiagram {
	diag := &diagrams.PlantUMLDiagram{
		Name:      "Folder Structure",
		Classes:   make([]diagrams.Class, 0),
		Relations: make([]diagrams.Relation, 0),
	}
	existingRels := make(map[string]bool)

	for _, ext := range c.Extensions {
		if len(ext.DocTypes) == 0 {
			// exclude components that are not schema related
			continue
		}
		for _, docType := range ext.DocTypes {
			if docType.IsInvisible() {
				continue
			}
			for _, subType := range docType.Subtypes.Types {
				searchKey := fmt.Sprintf("%s:%s", docType.Name, subType)
				if _, found := existingRels[searchKey]; found {
					continue
				} else {
					existingRels[searchKey] = true
				}
				rel := diagrams.Relation{
					From: docType.Name,
					To:   subType,
				}
				diag.Relations = append(diag.Relations, rel)
			}

		}
	}
	return diag
}

func GenerateDocumentInheritance(c Component) *diagrams.PlantUMLDiagram {
	diag := &diagrams.PlantUMLDiagram{
		Name:      "Custom Document Type Inheritance",
		Classes:   make([]diagrams.Class, 0),
		Relations: make([]diagrams.Relation, 0),
	}

	for _, ext := range c.Extensions {
		if len(ext.DocTypes) == 0 {
			// exclude components that are not schema related
			continue
		}
		for _, docType := range ext.DocTypes {
			if docType.IsInvisible() {
				continue
			}
			class := diagrams.Class{
				Type:    "class",
				Name:    docType.Name,
				Content: make([]diagrams.Writable, 0),
			}
			diag.Classes = append(diag.Classes, class)

			if docType.Extends != "" {
				rel := diagrams.Relation{
					From:  docType.Extends,
					To:    docType.Name,
					Label: "extends",
				}
				diag.Relations = append(diag.Relations, rel)
			}

			for _, subType := range docType.Subtypes.Types {
				field := diagrams.Field{
					Name: subType,
				}
				class.Content = append(class.Content, field)
			}

		}
	}

	return diag
}
