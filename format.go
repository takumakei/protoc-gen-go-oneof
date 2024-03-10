package main

import (
	"bufio"
	_ "embed"
	"io"
	"strings"
	"text/template"

	sprig "github.com/Masterminds/sprig/v3"
)

//go:embed format.tmpl
var formatTmpl string

var format = template.Must(
	template.New("").Funcs(sprig.TxtFuncMap()).Parse(formatTmpl),
)

func Formatf(w io.Writer, model *Model) (err error) {
	switch o := w.(type) {
	case *strings.Builder:

	case interface{ Flush() error }:
		defer func() {
			if err == nil {
				err = o.Flush()
			}
		}()

	default:
		b := bufio.NewWriter(o)
		defer func() {
			if err == nil {
				err = b.Flush()
			}
		}()
		w = b
	}
	return format.Execute(w, model)
}

func Format(model *Model) (string, error) {
	s := new(strings.Builder)
	err := Formatf(s, model)
	return s.String(), err
}

type Model struct {
	Versions Versions
	Source   string // e.g. "example/v1/hello.proto"
	Package  string // e.g. "examplev1"

	Oneofs []*Oneof
}

type Versions struct {
	ProtocGenGoOneof string
	Protoc           string
}

type Oneof struct {
	Name      string // e.g. "property"
	Parent    string // e.g. "Hello"
	FullName  string // e.g. "HelloProperty"
	Interface string // e.g. "isHello_Property"

	Fields []*Field
}

type Field struct {
	Name string // e.g. "email", "world", "create_time"
	Type string // e.g. "string", "*World", "*timestamppb.Timestamp"

	Struct string // e.g. "Hello_Email", "Hello_World", "Hello_CreateTime"
}
