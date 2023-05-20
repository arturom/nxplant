package structuretemplate

import "github.com/arturom/nxplant/diagrams"

func GenerateStructureTemplate(c Component) *diagrams.Diagram {
	diag := &diagrams.Diagram{
		Name:      "Structure Templates",
		Classes:   make([]diagrams.Class, 0),
		Relations: make([]diagrams.Relation, 0),
	}

	for _, ext := range c.Extensions {
		for _, factory := range ext.FactoryBindings {
			for _, templateItem := range factory.Template.TemplateItems {
				rel := diagrams.Relation{
					From: factory.TargetType,
					To:   templateItem.TypeName,
				}
				diag.Relations = append(diag.Relations, rel)
			}
		}
	}
	return diag
}
