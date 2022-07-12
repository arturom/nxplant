package nxplant

import (
	"fmt"
	"testing"
)

func TestRenderSchema(t *testing.T) {
	userSchema := Schema{
		Name:   "user",
		Prefix: "user",
		Fields: FieldSet{
			"name":  "String",
			"email": "String",
		},
	}
	result := RenderSchema(userSchema)
	fmt.Println(result)
}
