package invoke

import (
	"context"

	common "github.com/dapr/go-sdk/service/common"
)

// HandlerFunc ...
type HandlerFunc func(ctx context.Context, in *common.InvocationEvent) (*common.Content, error)

// HandlerDef ...
type HandlerDef struct {
	Name    string
	Handler interface{}
}
