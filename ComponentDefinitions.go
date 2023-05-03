package main

import (
	"encoding/xml"
)

type Component struct {
	XMLName    xml.Name    `xml:"component"`
	Name       string      `xml:"name,attr"`
	Extensions []Extension `xml:"extension"`
}

type Extension struct {
	XMLName  xml.Name           `xml:"extension"`
	Target   string             `xml:"target,attr"`
	Point    string             `xml:"point,attr"`
	Schemas  []ComponentSchema  `xml:"schema"`
	DocTypes []ComponentDocType `xml:"doctype"`
}

type ComponentSchema struct {
	XMLName xml.Name `xml:"schema"`
}

type ComponentDocType struct {
	XMLName  xml.Name `xml:"doctype"`
	Name     string   `xml:"name,attr"`
	Extends  string   `xml:"extends,attr"`
	Append   bool     `xml:"append,attr"`
	Subtypes SubTypes `xml:"subtypes"`
}

type SubTypes struct {
	Types []string `xml:"type"`
}

/*
func WriteFolderStructureDiagram(sb *strings.Builder, component Component) error {
	if _, err := sb.WriteString("@startuml docTypes\n\n"); err != nil {
		return err
	}
	for _, e := range component.Extensions {
		if len(e.DocTypes) == 0 {
			// exclude components that are not schema related
			continue
		}
		for _, d := range e.DocTypes {
			if _, err := sb.WriteString(fmt.Sprintf("class %s\n", d.Name)); err != nil {
				return err
			}
			if d.Extends != "" {

				if _, err := sb.WriteString(fmt.Sprintf("%s <|-- %s\n", d.Name, d.Extends)); err != nil {
					return err
				}
			}

			for _, s := range d.Subtypes.Types {
				if _, err := sb.WriteString(fmt.Sprintf("  - %s\n", s)); err != nil {
					return err
				}
			}
		}
	}
	if _, err := sb.WriteString("@enduml"); err != nil {
		return err
	}
	return nil
}

func WriteDocumentHierarchy(sb *strings.Builder, component Component) error {
	if _, err := sb.WriteString("@startuml studioDocTypes\n"); err != nil {
		return err
	}
	for _, e := range component.Extensions {
		if len(e.DocTypes) == 0 {
			continue
		}
		for _, d := range e.DocTypes {
			if len(d.Subtypes.Types) == 0 {
				if _, err := sb.WriteString(fmt.Sprintf("\nclass %s\n", d.Name)); err != nil {
					return err
				}
			} else {
				if _, err := sb.WriteString(fmt.Sprintf("\nclass %s {\n", d.Name)); err != nil {
					return err
				}

				existing := make(map[string]bool)
				for _, s := range d.Subtypes.Types {
					_, exists := existing[s]
					if exists {
						continue
					}
					if _, err := sb.WriteString(fmt.Sprintf("  -%s\n", s)); err != nil {
						return err
					}
				}
				if _, err := sb.WriteString("}\n"); err != nil {
					return err
				}
			}

			if d.Extends != "" {
				if _, err := sb.WriteString(fmt.Sprintf("%s <|-- %s\n", d.Extends, d.Name)); err != nil {
					return err
				}
			}
		}
	}
	if _, err := sb.WriteString("@enduml"); err != nil {
		return err
	}
	return nil
}
*/
