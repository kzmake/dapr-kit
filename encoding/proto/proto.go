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
	registry.Encodings[ContentTypeApplicationXProtobuf] = NewEncoding()
	registry.Encodings[ContentTypeApplocationGRPC] = NewEncoding()
}

type protoEncoding struct{}

// NewEncoding ...
func NewEncoding() kit.Encoding {
	return &protoEncoding{}
}

func (c *protoEncoding) Marshal(v interface{}) ([]byte, error) {
	m, ok := v.(proto.Message)
	if !ok {
		return nil, xerrors.Errorf("failed to cast proto.Message")
	}

	return proto.Marshal(m)
}

func (c *protoEncoding) Unmarshal(data []byte, v interface{}) error {
	m, ok := v.(proto.Message)
	if !ok {
		return xerrors.Errorf("failed to cast proto.Message")
	}

	return proto.Unmarshal(data, m)
}
