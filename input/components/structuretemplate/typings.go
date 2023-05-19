package structuretemplate

import "encoding/xml"

type Component struct {
	XMLName    xml.Name    `xml:"component"`
	Name       string      `xml:"name,attr"`
	Extensions []Extension `xml:"extension"`
}

type Extension struct {
	XMLName         xml.Name         `xml:"extension"`
	Target          string           `xml:"target,attr"`
	Point           string           `xml:"point,attr"`
	FactoryBindings []FactoryBinding `xml:"factoryBinding"`
}

type FactoryBinding struct {
	Name        string   `xml:"name,attr"`
	FactoryName string   `xml:"factoryName,attr"`
	TargetType  string   `xml:"targetType,attr"`
	Template    Template `xml:"template"`
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
