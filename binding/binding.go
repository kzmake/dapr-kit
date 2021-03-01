package invoke

import (
	"context"

	common "github.com/dapr/go-sdk/service/common"
)

// Func ...
type Func func(ctx context.Context, in *common.BindingEvent) ([]byte, error)

// Def ...
type Def struct {
	Name    string
	Handler interface{}
}
