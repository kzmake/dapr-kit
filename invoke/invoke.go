package invoke

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	common "github.com/dapr/go-sdk/service/common"
	"google.golang.org/protobuf/proto"
)

// HandlerFunc ...
type HandlerFunc func(ctx context.Context, in *common.InvocationEvent) (*common.Content, error)

// HandlerDef ...
type HandlerDef struct {
	Name    string
	Handler interface{}
}

// Decode ...
func Decode(in *common.InvocationEvent, req proto.Message) error {
	if len(in.Data) == 0 {
		return nil
	}

	switch strings.TrimSpace(strings.Split(in.ContentType, ";")[0]) {
	case "application/json":
		return json.Unmarshal(in.Data, req)
	case "application/x-protobuf":
		return proto.Unmarshal(in.Data, req)
	default:
		return fmt.Errorf("unknown content type")
	}
}

// Encode ...
func Encode(res proto.Message) (*common.Content, error) {
	d, err := proto.Marshal(res)
	if err != nil {
		return nil, err
	}

	return &common.Content{
		ContentType: "application/x-protobuf",
		Data:        d,
	}, nil
}
