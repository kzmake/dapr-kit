package util

import (
	"context"
	"fmt"
	"reflect"

	common "github.com/dapr/go-sdk/service/common"
)

// RegisterInvocationHandlers ...
func RegisterInvocationHandlers(s common.Service, impl interface{}, serviceName string) error {
	typ := reflect.TypeOf(impl)
	val := reflect.ValueOf(impl)

	for i := 0; i < typ.NumMethod(); i++ {
		method := val.Method(i)
		methodName := typ.Method(i).Name

		callNames := []string{
			methodName,
			serviceName + "." + methodName,
			serviceName + "/" + methodName,
		}

		for _, callName := range callNames {
			if err := s.AddServiceInvocationHandler(callName, newInvocationHandler(method)); err != nil {
				return err
			}
		}
	}

	return nil
}

// newInvocationHandler ...
func newInvocationHandler(method reflect.Value) func(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
	reqType := method.Type().In(1) // get req type

	return func(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
		if in.ContentType != ContentType {
			return nil, fmt.Errorf("cont")
		}

		req := reflect.ValueOf(reflect.New(reqType).Interface()).Elem().Interface()

		if err := Unmarshal(in.Data, &req); err != nil {
			return nil, err
		}

		result := method.Call([]reflect.Value{
			reflect.ValueOf(ctx),
			reflect.ValueOf(req),
		})
		if !result[1].IsNil() {
			if err := result[1].Interface().(error); err != nil {
				return nil, err
			}
		}

		res := result[0].Interface()

		v, err := Marshal(res)
		if err != nil {
			return nil, err
		}

		return &common.Content{
			ContentType: ContentType,
			Data:        v,
		}, nil
	}
}
