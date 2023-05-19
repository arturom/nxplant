package structuretemplate

import "encoding/xml"

type StructureComponent struct {
	XMLName    xml.Name             `xml:"component"`
	Name       string               `xml:"name,attr"`
	Extensions []StructureExtension `xml:"extension"`
}

type StructureExtension struct {
	XMLName        xml.Name       `xml:"extension"`
	Target         string         `xml:"target,attr"`
	Point          string         `xml:"point,attr"`
	FactoryBinding FactoryBinding `xml:"factoryBinding"`
}

type FactoryBinding struct {
	Template Template `xml:"template"`
}

type Template struct {
	TemplateItems []TemplateItem `xml:"templateItem"`
}

type TemplateItem struct {
	TypeName    string `xml:"typeName,attr"`
	ID          string `xml:"id,attr"`
	Title       string `xml:"title,attr"`
	Description string `xml:"description,attr"`
}