package json

import (
	jsoniter "github.com/json-iterator/go"

	kit "github.com/kzmake/dapr-kit"
	"github.com/kzmake/dapr-kit/registry"
)

// ContentType ...
const (
	ContentTypeApplicationJSON = "application/json"
)

func init() {
	registry.Encodings[ContentTypeApplicationJSON] = NewEncoding()
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type jsonEncoding struct{}

// NewEncoding ...
func NewEncoding() kit.Encoding {
	return &jsonEncoding{}
}

func (c *jsonEncoding) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *jsonEncoding) Unmarshal(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal(data, v)
}
