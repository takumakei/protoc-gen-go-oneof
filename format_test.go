package main

import (
	_ "embed"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed format.example
var want string

func TestFormat(t *testing.T) {
	model := &Model{
		Versions: Versions{
			ProtocGenGoOneof: "1.2.3-p0+b1",
			Protoc:           "1.23.456",
		},
		Source:  "example/v1/hello.proto",
		Package: "examplev1",

		Oneofs: []*Oneof{
			{
				Name:      "property",
				Parent:    "Hello",
				FullName:  "HelloProperty",
				Interface: "isHello_Property",

				Fields: []*Field{
					{
						Name:   "number",
						Type:   "int32",
						Struct: "isHello_Number",
					},
					{
						Name:   "world",
						Type:   "*World",
						Struct: "isHello_World",
					},
				},
			},
		},
	}
	s, err := Format(model)
	require.NoError(t, err)
	assert.Equal(t, want, s)

	if yes, _ := strconv.ParseBool(os.Getenv("FORMAT_EXAMPLE")); yes {
		require.NoError(t, os.WriteFile("format.example", []byte(s), 0644))
	}
}
