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

type NamedItem struct {
	Name string `xml:"name,attr"`
}

type ComponentDocType struct {
	XMLName  xml.Name    `xml:"doctype"`
	Name     string      `xml:"name,attr"`
	Extends  string      `xml:"extends,attr"`
	Append   bool        `xml:"append,attr"`
	Subtypes SubTypes    `xml:"subtypes"`
	Schemas  []NamedItem `xml:"schema"`
	Facets   []NamedItem `xml:"facet"`
}

func (t ComponentDocType) isInvisible() bool {
	return t.containsFacet("HiddenInNavigation")
}

func (t ComponentDocType) containsFacet(facet string) bool {
	for _, f := range t.Facets {
		if f.Name == facet {
			return true
		}
	}
	return false
}

type SubTypes struct {
	Types []string `xml:"type"`
}

func (s SubTypes) containsSubType(subtype string) bool {
	for _, t := range s.Types {
		if t == subtype {
			return true
		}
	}
	return false
}
