package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestRenderSchema(t *testing.T) {
	userSchema := RestSchema{
		Name:   "user",
		Prefix: "user",
		Fields: FieldSet{
			"name":  "String",
			"email": "String",
		},
	}
	sb := &strings.Builder{}
	err := RenderSchema(sb, userSchema)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sb.String())
}
