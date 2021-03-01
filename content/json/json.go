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
	registry.Binders[ContentTypeApplicationJSON] = NewBinder()
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type jsonBinder struct{}

// NewBinder ...
func NewBinder() kit.Binder {
	return &jsonBinder{}
}

func (c *jsonBinder) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *jsonBinder) Unmarshal(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal(data, v)
}
