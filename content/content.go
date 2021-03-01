package content

import (
	"fmt"
	"strings"

	"golang.org/x/xerrors"

	kit "github.com/kzmake/dapr-kit"
	"github.com/kzmake/dapr-kit/registry"

	// supported content types
	_ "github.com/kzmake/dapr-kit/content/json"
	_ "github.com/kzmake/dapr-kit/content/proto"
)

var (
	// NewBinder ...
	NewBinder = newBinder
)

// newBinder ...
func newBinder(contentType string) (kit.Binder, error) {
	ct := strings.TrimSpace(strings.Split(contentType, ";")[0])
	fmt.Println(ct)
	if b, ok := registry.Binders[ct]; ok {
		return b, nil
	}

	return nil, xerrors.Errorf("unsupported content-type: %s", ct)
}
