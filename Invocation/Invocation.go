package invocation

import (
	"context"

	common "github.com/dapr/go-sdk/service/common"
)

// Func ...
type Func func(ctx context.Context, in *common.InvocationEvent) (*common.Content, error)

// Def ...
type Def struct {
	Name    string
	Handler interface{}
}
