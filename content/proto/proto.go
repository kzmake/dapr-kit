package proto

import (
	"golang.org/x/xerrors"

	"google.golang.org/protobuf/proto"

	kit "github.com/kzmake/dapr-kit"
	"github.com/kzmake/dapr-kit/registry"
)

// ContentType ...
const (
	ContentTypeApplicationXProtobuf = "application/x-protobuf"
	ContentTypeApplocationGRPC      = "application/grpc"
)

func init() {
	registry.Binders[ContentTypeApplicationXProtobuf] = NewBinder()
	registry.Binders[ContentTypeApplocationGRPC] = NewBinder()
}

type protoBinder struct{}

// NewBinder ...
func NewBinder() kit.Binder {
	return &protoBinder{}
}

func (c *protoBinder) Marshal(v interface{}) ([]byte, error) {
	m, ok := v.(proto.Message)
	if !ok {
		return nil, xerrors.Errorf("failed to cast proto.Message")
	}

	return proto.Marshal(m)
}

func (c *protoBinder) Unmarshal(data []byte, v interface{}) error {
	m, ok := v.(proto.Message)
	if !ok {
		return xerrors.Errorf("failed to cast proto.Message")
	}

	return proto.Unmarshal(data, m)
}
