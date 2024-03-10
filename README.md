protoc-gen-go-oneof
======================================================================

Generating golang helper functions for `oneof` fields of protocol buffers.

### example input

```proto
syntax = "proto3";

package example.v1;

message Hello {
  oneof property {
    int32 number = 1;
    World world = 2;
  }
}

message World {
  string name = 1;
}
```

### example code generated by protoc-gen-go-oneof

```go
// Code generated by protoc-gen-go-oneof. DO NOT EDIT.
// versions:
// 	protoc-gen-go-oneof (devel)
// 	protoc              (unknown)
// source: example/v1/hello.proto

package examplev1

// OneofHelloProperty can contain one of the value for Hello.property.
//
// Use the constructor for each oneof fields.
//
//   - func OneofHelloProperty_Number(number int32) OneofHelloProperty
//   - func OneofHelloProperty_World(world *World) OneofHelloProperty
type OneofHelloProperty struct {
	oneof isHello_Property
}

// Get returns the value for Hello.property.
func (property OneofHelloProperty) Get() isHello_Property {
	return property.oneof
}

// OneofHelloPropertyFrom returns OneofHelloProperty from the value of isHello_Property.
func OneofHelloPropertyFrom(property isHello_Property) OneofHelloProperty {
	return OneofHelloProperty{oneof: property}
}

// OneofHelloProperty_Number returns OneofHelloProperty.
func OneofHelloProperty_Number(number int32) OneofHelloProperty {
	return OneofHelloProperty{oneof: &Hello_Number{Number: number}}
}

// OneofHelloProperty_World returns OneofHelloProperty.
func OneofHelloProperty_World(world *World) OneofHelloProperty {
	return OneofHelloProperty{oneof: &Hello_World{World: world}}
}

type OneofHelloPropertyKind string

const (
	OneofHelloPropertyIsUnknown OneofHelloPropertyKind = "unknown"
	OneofHelloPropertyIsNil     OneofHelloPropertyKind = "nil"
	OneofHelloPropertyIsNumber  OneofHelloPropertyKind = "number"
	OneofHelloPropertyIsWorld   OneofHelloPropertyKind = "world"
)

func OneofHelloPropertyKindOf(property isHello_Property) OneofHelloPropertyKind {
	if property == nil {
		return OneofHelloPropertyIsNil
	}
	switch property.(type) {
	case *Hello_Number:
		return OneofHelloPropertyIsNumber
	case *Hello_World:
		return OneofHelloPropertyIsWorld
	}
	return OneofHelloPropertyIsUnknown
}
```

### usage example of helper functions

```go
func GetProperty() examplev1.OneofHelloProperty {
	if number, err := strconv.Atoi(os.Getenv("NUMBER")); err == nil {
		return examplev1.OneofHelloProperty_Number(int32(number))
	}
	if world := strings.TrimSpace(os.Getenv("WORLD")); world != "" {
		return examplev1.OneofHelloProperty_World(&examplev1.World{Name: world})
	}
	return examplev1.OneofHelloProperty{}
}

func New() *examplev1.Hello {
	return &examplev1.Hello{
		Property: GetProperty().Get(),
	}
}
```