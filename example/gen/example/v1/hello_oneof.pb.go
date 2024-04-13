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
//   - func OneofHelloProperty_Number(_number int32) OneofHelloProperty
//   - func OneofHelloProperty_World(_world *World) OneofHelloProperty
//   - func OneofHelloProperty_String(_string string) OneofHelloProperty
//   - func OneofHelloProperty_Type(_type string) OneofHelloProperty
//   - func OneofHelloProperty_Ratio(_ratio float32) OneofHelloProperty
type OneofHelloProperty struct {
	oneof isHello_Property
}

// Get returns the value for Hello.property.
func (_property OneofHelloProperty) Get() isHello_Property {
	return _property.oneof
}

// OneofHelloPropertyFrom returns OneofHelloProperty from the value of isHello_Property.
func OneofHelloPropertyFrom(_property isHello_Property) OneofHelloProperty {
	return OneofHelloProperty{oneof: _property}
}

// OneofHelloProperty_Number returns OneofHelloProperty.
func OneofHelloProperty_Number(_number int32) OneofHelloProperty {
	return OneofHelloProperty{oneof: &Hello_Number{Number: _number}}
}

// OneofHelloProperty_World returns OneofHelloProperty.
func OneofHelloProperty_World(_world *World) OneofHelloProperty {
	return OneofHelloProperty{oneof: &Hello_World{World: _world}}
}

// OneofHelloProperty_String returns OneofHelloProperty.
func OneofHelloProperty_String(_string string) OneofHelloProperty {
	return OneofHelloProperty{oneof: &Hello_String_{String_: _string}}
}

// OneofHelloProperty_Type returns OneofHelloProperty.
func OneofHelloProperty_Type(_type string) OneofHelloProperty {
	return OneofHelloProperty{oneof: &Hello_Type{Type: _type}}
}

// OneofHelloProperty_Ratio returns OneofHelloProperty.
func OneofHelloProperty_Ratio(_ratio float32) OneofHelloProperty {
	return OneofHelloProperty{oneof: &Hello_Ratio{Ratio: _ratio}}
}

type OneofHelloPropertyKind string

const (
	OneofHelloPropertyIsUnknown OneofHelloPropertyKind = "unknown"
	OneofHelloPropertyIsNil     OneofHelloPropertyKind = "nil"
	OneofHelloPropertyIsNumber  OneofHelloPropertyKind = "number"
	OneofHelloPropertyIsWorld   OneofHelloPropertyKind = "world"
	OneofHelloPropertyIsString  OneofHelloPropertyKind = "string"
	OneofHelloPropertyIsType    OneofHelloPropertyKind = "type"
	OneofHelloPropertyIsRatio   OneofHelloPropertyKind = "ratio"
)

func OneofHelloPropertyKindOf(_property isHello_Property) OneofHelloPropertyKind {
	if _property == nil {
		return OneofHelloPropertyIsNil
	}
	switch _property.(type) {
	case *Hello_Number:
		return OneofHelloPropertyIsNumber
	case *Hello_World:
		return OneofHelloPropertyIsWorld
	case *Hello_String_:
		return OneofHelloPropertyIsString
	case *Hello_Type:
		return OneofHelloPropertyIsType
	case *Hello_Ratio:
		return OneofHelloPropertyIsRatio
	}
	return OneofHelloPropertyIsUnknown
}
