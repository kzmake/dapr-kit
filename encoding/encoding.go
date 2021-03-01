package encoding

import (
	"fmt"
	"strings"

	"golang.org/x/xerrors"

	kit "github.com/kzmake/dapr-kit"
	"github.com/kzmake/dapr-kit/registry"

	// supported content types
	_ "github.com/kzmake/dapr-kit/encoding/json"
	_ "github.com/kzmake/dapr-kit/encoding/proto"
)

var (
	// New ...
	New = new
)

// new ...
func new(contentType string) (kit.Encoding, error) {
	ct := strings.TrimSpace(strings.Split(contentType, ";")[0])
	fmt.Println(ct)
	if b, ok := registry.Encodings[ct]; ok {
		return b, nil
	}

	return nil, xerrors.Errorf("unsupported content-type: %s", ct)
}
