package example

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	examplev1 "github.com/takumakei/protoc-gen-go-oneof/test/gen/example/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestOneOfHelloProperty(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var actu examplev1.OneofHelloProperty
		var want examplev1.Hello
		assert.Equal(t, want.Property, actu.Get())
	})

	t.Run("ctor", func(t *testing.T) {
		var v examplev1.OneofHelloProperty
		assert.Nil(t, v.Get())

		j, err := protojson.Marshal(&examplev1.Hello{Property: v.Get()})
		assert.NoError(t, err)
		w := `{}`
		assert.Equal(t, w, string(j))
		assert.Equal(t, examplev1.OneofHelloPropertyIsNil, examplev1.OneofHelloPropertyKindOf(v.Get()))

		v = examplev1.OneofHelloProperty_Email("EMAIL")
		j, err = protojson.Marshal(&examplev1.Hello{Property: v.Get()})
		assert.NoError(t, err)
		w = `{"email":"EMAIL"}`
		assert.Equal(t, w, string(j))
		assert.Equal(t, examplev1.OneofHelloPropertyIsEmail, examplev1.OneofHelloPropertyKindOf(v.Get()))

		v = examplev1.OneofHelloProperty_Number(42)
		j, err = protojson.Marshal(&examplev1.Hello{Property: v.Get()})
		assert.NoError(t, err)
		w = `{"number":42}`
		assert.Equal(t, w, string(j))
		assert.Equal(t, examplev1.OneofHelloPropertyIsNumber, examplev1.OneofHelloPropertyKindOf(v.Get()))

		v = examplev1.OneofHelloProperty_World(&examplev1.World{Name: "WORLD"})
		j, err = protojson.Marshal(&examplev1.Hello{Property: v.Get()})
		assert.NoError(t, err)
		w = `{"world":{"name":"WORLD"}}`
		assert.Equal(t, w, string(j))
		assert.Equal(t, examplev1.OneofHelloPropertyIsWorld, examplev1.OneofHelloPropertyKindOf(v.Get()))

		v = examplev1.OneofHelloProperty_CreateTime(timestamppb.New(time.Unix(123456789, 123456789)))
		j, err = protojson.Marshal(&examplev1.Hello{Property: v.Get()})
		assert.NoError(t, err)
		w = `{"createTime":"1973-11-29T21:33:09.123456789Z"}`
		assert.Equal(t, w, string(j))
		assert.Equal(t, examplev1.OneofHelloPropertyIsCreateTime, examplev1.OneofHelloPropertyKindOf(v.Get()))
	})
}
