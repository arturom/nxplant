package osgi

import (
	"encoding/xml"
	"fmt"
)

type Component struct {
	XMLName    xml.Name    `xml:"component"`
	Name       string      `xml:"name,attr"`
	Extensions []Extension `xml:"extension"`
}

type Extension struct {
	XMLName  xml.Name  `xml:"extension"`
	Target   string    `xml:"target,attr"`
	Point    string    `xml:"point,attr"`
	Schemas  []Schema  `xml:"schema"`
	DocTypes []DocType `xml:"doctype"`
}

type Schema struct {
	XMLName xml.Name `xml:"schema"`
}

type DocType struct {
	XMLName  xml.Name `xml:"doctype"`
	Name     string   `xml:"name,attr"`
	Extends  string   `xml:"extends,attr"`
	Append   bool     `xml:"append,attr"`
	Subtypes SubTypes `xml:"subtypes"`
}

type SubTypes struct {
	Types []string `xml:"type"`
}

func ParseXML(filepath string) Component {
	component := Component{}
	str := "abc"
	x := []byte(str)
	xml.Unmarshal(x, &component)
	return component
}

func GenerateFolderStructure(component Component) {
	for _, e := range component.Extensions {
		if len(e.DocTypes) == 0 {
			continue
		}
		fmt.Printf("extension: %s %s \n", e.Target, e.Point)
		for _, d := range e.DocTypes {
			// fmt.Printf("- %s / %s / %v\n", d.Name, d.Extends, d.Append)
			fmt.Printf("class %s\n", d.Name)
			if d.Extends != "" {
				fmt.Printf("%s <|-- %s\n", d.Name, d.Extends)
			}
			for _, s := range d.Subtypes.Types {
				fmt.Printf("  - %s\n", s)
			}
		}
		fmt.Println()
	}
}

func GenerateHierarchy(component Component) {
	fmt.Println(component.Name)
	for _, e := range component.Extensions {
		if len(e.DocTypes) == 0 {
			continue
		}
		fmt.Printf("extension: %s %s \n", e.Target, e.Point)
		for _, d := range e.DocTypes {
			// fmt.Printf("- %s / %s / %v\n", d.Name, d.Extends, d.Append)
			fmt.Printf("class %s\n", d.Name)
			if d.Extends != "" {
				fmt.Printf("%s <|-- %s\n", d.Name, d.Extends)
			}
			for _, s := range d.Subtypes.Types {
				fmt.Printf("  - %s\n", s)
			}
		}
		fmt.Println()
	}
}
