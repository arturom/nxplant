package main

import (
	"encoding/xml"
	"fmt"
	"strings"
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

func ParseXML(filepath string) Component {
	component := Component{}
	str := "abc"
	x := []byte(str)
	xml.Unmarshal(x, &component)
	return component
}

func GenerateFolderStructure(sb *strings.Builder, component Component) error {
	for _, e := range component.Extensions {
		if len(e.DocTypes) == 0 {
			// exclude components that are not schema related
			continue
		}
		if _, err := sb.WriteString(fmt.Sprintf("extension: %s %s \n", e.Target, e.Point)); err != nil {
			return err
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
	return nil
}

func GenerateHierarchy(sb *strings.Builder, component Component) error {
	if _, err := sb.WriteString(component.Name); err != nil {
		return err
	}
	for _, e := range component.Extensions {
		if len(e.DocTypes) == 0 {
			continue
		}
		if _, err := sb.WriteString(fmt.Sprintf("extension: %s %s \n", e.Target, e.Point)); err != nil {
			return err
		}
		for _, d := range e.DocTypes {
			// fmt.Printf("- %s / %s / %v\n", d.Name, d.Extends, d.Append)
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
	return nil
}
