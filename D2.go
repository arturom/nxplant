package main

import (
	"fmt"
	"strings"
)

func writeField(f Field, sb *strings.Builder) error {
	if f.Type == "" {
		if _, err := sb.WriteString(fmt.Sprintf("  %s %s\n", f.Char, f.Name)); err != nil {
			return err
		}
	} else {
		if _, err := sb.WriteString(fmt.Sprintf("  %s%s: %s\n", f.Char, f.Name, f.Type)); err != nil {
			return err
		}
	}
	return nil
}

func writeWritable(w Writable, sb *strings.Builder) error {
	switch w.(type) {
	case Field:
		if err := writeField(w.(Field), sb); err != nil {
			return err
		}
	}
	return nil
}

func writeClass(class Class, sb *strings.Builder) error {
	if _, err := sb.WriteString(fmt.Sprintf("%s: {\n", class.Name)); err != nil {
		return err
	}
	if _, err := sb.WriteString("  shape: class\n\n"); err != nil {
		return err
	}

	for _, w := range class.content {
		if err := writeWritable(w, sb); err != nil {
			return err
		}
	}

	if _, err := sb.WriteString("}\n\n\n"); err != nil {
		return err
	}
	return nil
}

func writeRelationship(rel Relation, sb *strings.Builder) error {
	if _, err := sb.WriteString(fmt.Sprintf("%s -> %s\n", rel.From, rel.To)); err != nil {
		return err
	}
	return nil
}

func WriteD2(diagram PlantUMLDiagram, sb *strings.Builder) error {
	for _, class := range diagram.Classes {
		if err := writeClass(class, sb); err != nil {
			return err
		}
	}
	for _, rel := range diagram.Relations {
		if err := writeRelationship(rel, sb); err != nil {
			return err
		}
	}
	return nil
}
