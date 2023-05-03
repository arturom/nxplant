package main

import (
	"fmt"
	"strings"
)

type Writable interface {
	write(sb *strings.Builder) error
}

type Field struct {
	Char string
	Type string
	Name string
}

func (f Field) write(sb *strings.Builder) error {
	char := f.Char
	if char == "" {
		char = "-"
	}
	if _, err := sb.WriteString(fmt.Sprintf("  %s%s %s\n", char, f.Type, f.Name)); err != nil {
		return err
	}
	return nil
}

type Separation struct {
	val string
}

func (s Separation) write(sb *strings.Builder) error {
	if _, err := sb.WriteString(fmt.Sprintf("-- %s --\n", s.val)); err != nil {
		return err
	}
	return nil
}

type Class struct {
	Type    string
	Name    string
	content []Writable
}

func (c *Class) hasFields() bool {
	return len(c.content) > 0
}

func (c *Class) write(sb *strings.Builder) error {
	type_ := c.Type
	if type_ == "" {
		type_ = "class"
	}
	if _, err := sb.WriteString(fmt.Sprintf("%s \"%s\"", c.Type, c.Name)); err != nil {
		return err
	}
	if c.hasFields() {
		if _, err := sb.WriteString(" {\n"); err != nil {
			return err
		}
		for _, d := range c.content {
			if err := d.write(sb); err != nil {
				return err
			}
		}
		if _, err := sb.WriteString("}\n\n"); err != nil {
			return err
		}
	} else {
		if _, err := sb.WriteString("\n\n"); err != nil {
			return err
		}

	}
	return nil
}

type Relation struct {
	From   string
	To     string
	Label  string
	Symbol string
}

func (r *Relation) write(sb *strings.Builder) error {
	symbol := r.Symbol
	if symbol == "" {
		symbol = "<|--"
	}
	if _, err := sb.WriteString(fmt.Sprintf("\"%s\" %s \"%s\"\n", r.From, symbol, r.To)); err != nil {
		return err
	}
	return nil
}

type Diagram struct {
	Name      string
	Classes   []Class
	Relations []Relation
}

func (d *Diagram) write(sb *strings.Builder) error {
	if _, err := sb.WriteString(fmt.Sprintf("@startuml \"%s\" \n\n", d.Name)); err != nil {
		return err
	}

	if _, err := sb.WriteString(fmt.Sprintln("title", d.Name, "\n")); err != nil {
		return err
	}

	/*
		if _, err := sb.WriteString("!theme spacelab\n"); err != nil {
			return err
		}
	*/

	for _, class := range d.Classes {
		if err := class.write(sb); err != nil {
			return err
		}
	}

	for _, rel := range d.Relations {
		if err := rel.write(sb); err != nil {
			return err
		}
	}

	if _, err := sb.WriteString("\n@enduml"); err != nil {
		return err
	}
	return nil
}
