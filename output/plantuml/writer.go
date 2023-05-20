package plantuml

import (
	"fmt"
	"strings"

	"github.com/arturom/nxplant/diagrams"
)

func writeField(f diagrams.Field, sb *strings.Builder) error {
	char := f.Char
	if char == "" {
		char = "-"
	}
	if _, err := sb.WriteString(fmt.Sprintf("  %s%s %s\n", char, f.Type, f.Name)); err != nil {
		return err
	}
	return nil
}

func writeSeparation(s diagrams.Separation, sb *strings.Builder) error {
	if _, err := sb.WriteString(fmt.Sprintf("-- %s --\n", s.Val)); err != nil {
		return err
	}
	return nil
}

func writeClass(c diagrams.Class, sb *strings.Builder) error {
	type_ := c.Type
	if type_ == "" {
		type_ = "class"
	}
	if _, err := sb.WriteString(fmt.Sprintf("%s \"%s\"", c.Type, c.Name)); err != nil {
		return err
	}
	if c.HasFields() {
		if _, err := sb.WriteString(" {\n"); err != nil {
			return err
		}
		for _, d := range c.Content {
			switch v := d.(type) {
			case diagrams.Field:
				if err := writeField(v, sb); err != nil {
					return err
				}
			case diagrams.Separation:
				if err := writeSeparation(v, sb); err != nil {
					return err
				}
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

func writeRelation(r diagrams.Relation, sb *strings.Builder) error {
	symbol := r.Symbol
	if symbol == "" {
		symbol = "<|--"
	}
	if _, err := sb.WriteString(fmt.Sprintf("\"%s\" %s \"%s\"\n", r.From, symbol, r.To)); err != nil {
		return err
	}
	return nil
}

func WritePlantuml(d diagrams.Diagram, sb *strings.Builder) error {
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
		if err := writeClass(class, sb); err != nil {
			return err
		}
	}

	for _, rel := range d.Relations {
		if err := writeRelation(rel, sb); err != nil {
			return err
		}
	}

	if _, err := sb.WriteString("\n@enduml"); err != nil {
		return err
	}
	return nil
}
