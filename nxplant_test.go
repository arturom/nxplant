package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/arturom/nxplant/input/restapi/doctypes"
)

func TestRenderSchema(t *testing.T) {
	userSchema := doctypes.RestSchema{
		Name:   "user",
		Prefix: "user",
		Fields: doctypes.FieldSet{
			"name":  "String",
			"email": "String",
		},
	}
	sb := &strings.Builder{}
	err := doctypes.RenderSchema(sb, userSchema)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sb.String())
}
