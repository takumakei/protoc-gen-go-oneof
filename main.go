package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	var flagTemplate string
	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.StringVar(&flagTemplate, "template", flagTemplate, "template")

	protogen.Options{ParamFunc: flags.Set}.Run(func(gen *protogen.Plugin) error {
		if err := InitFormat(flagTemplate); err != nil {
			return err
		}

		// https://github.com/protocolbuffers/protobuf/blob/main/docs/implementing_proto3_presence.md
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		for _, protoFile := range gen.Files {
			if protoFile.Generate {
				if err := Generate(gen, protoFile); err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				}
			}
		}
		return nil
	})
}

func Generate(gen *protogen.Plugin, protoFile *protogen.File) error {
	if !HasOneof(protoFile.Messages) {
		return nil
	}

	filename := protoFile.GeneratedFilenamePrefix + "_oneof.pb.go"
	g := gen.NewGeneratedFile(filename, protoFile.GoImportPath)

	code, err := Format(NewModel(gen, protoFile, g))
	if err != nil {
		return err
	}
	g.P(code)
	return nil
}

func HasOneof(messages []*protogen.Message) bool {
	for _, m := range messages {
		if hasOneof(m) {
			return true
		}
	}
	return false
}

func hasOneof(m *protogen.Message) bool {
	for _, o := range m.Oneofs {
		if !o.Desc.IsSynthetic() {
			return true
		}
	}
	return HasOneof(m.Messages)
}

func NewModel(gen *protogen.Plugin, protoFile *protogen.File, g *protogen.GeneratedFile) *Model {
	return &Model{
		Versions: Versions{
			ProtocGenGoOneof: version,
			Protoc:           VersionString(gen.Request.CompilerVersion),
		},
		Source:  protoFile.GeneratedFilenamePrefix + ".proto",
		Package: string(protoFile.GoPackageName),
		Oneofs:  NewOneofs(gen, protoFile, g, protoFile.Messages),
	}
}

func VersionString(v *pluginpb.Version) string {
	if v == nil || (v.Major == nil && v.Minor == nil && v.Patch == nil && v.Suffix == nil) {
		return "(unknown)"
	}
	var (
		major  = v.GetMajor()
		minor  = v.GetMinor()
		patch  = v.GetPatch()
		suffix = v.GetSuffix()
	)
	return fmt.Sprintf("%d.%d.%d%s", major, minor, patch, suffix)
}

func NewOneofs(gen *protogen.Plugin, protoFile *protogen.File, g *protogen.GeneratedFile, messages []*protogen.Message) []*Oneof {
	var list []*Oneof
	for _, m := range messages {
		list = append(list, NewOneofsMessage(gen, protoFile, g, m)...)
		list = append(list, NewOneofs(gen, protoFile, g, m.Messages)...)
	}
	return list
}

func NewOneofsMessage(gen *protogen.Plugin, protoFile *protogen.File, g *protogen.GeneratedFile, m *protogen.Message) []*Oneof {
	var list []*Oneof
	for _, o := range m.Oneofs {
		if o.Desc.IsSynthetic() { // excludes optional
			continue
		}
		oneof := &Oneof{
			Name:      string(o.Desc.Name()),
			Parent:    m.GoIdent.GoName,
			FullName:  strings.ReplaceAll(o.GoIdent.GoName, "_", ""),
			Interface: "is" + o.GoIdent.GoName,
			Fields:    NewFields(gen, protoFile, g, m, o),
		}
		list = append(list, oneof)
	}
	return list
}

func NewFields(gen *protogen.Plugin, protoFile *protogen.File, g *protogen.GeneratedFile, m *protogen.Message, o *protogen.Oneof) []*Field {
	var list []*Field
	for _, f := range o.Fields {
		field := &Field{
			Name:   string(f.Desc.Name()),
			GoName: f.GoName,
			Type:   TypeOf(g, f),
			Struct: f.GoIdent.GoName,
		}
		list = append(list, field)
	}
	return list
}

func TypeOf(g *protogen.GeneratedFile, f *protogen.Field) string {
	if f.Desc.Kind() == protoreflect.MessageKind {
		return "*" + g.QualifiedGoIdent(f.Message.GoIdent)
	}
	if f.Desc.Kind() == protoreflect.EnumKind {
		return g.QualifiedGoIdent(f.Enum.GoIdent)
	}
	ident := f.Desc.Kind().String()
	switch ident {
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "sint32":
		return "int32"
	case "sint64":
		return "int64"
	case "fixed32":
		return "uint32"
	case "fixed64":
		return "uint64"
	case "sfixed32":
		return "int32"
	case "sfixed64":
		return "int64"
	}
	return ident
}
