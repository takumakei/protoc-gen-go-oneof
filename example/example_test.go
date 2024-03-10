package example

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	examplev1 "github.com/takumakei/protoc-gen-go-oneof/example/gen/example/v1"
)

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

func TestNew(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		t.Setenv("NUMBER", "42")
		j, err := json.Marshal(New())
		require.NoError(t, err)
		want := `{"Property":{"Number":42}}`
		assert.Equal(t, want, string(j))
	})

	t.Run("world", func(t *testing.T) {
		t.Setenv("WORLD", "Golang")
		j, err := json.Marshal(New())
		require.NoError(t, err)
		want := `{"Property":{"World":{"name":"Golang"}}}`
		assert.Equal(t, want, string(j))
	})
}
