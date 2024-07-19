package basic

import (
	"context"
	"encoding/json"
	"reflect"
)

func ContextWithValues(parent context.Context, vals map[string]any) context.Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if len(vals) == 0 {
		panic("empty map")
	}
	return &valuesCtx{parent, vals}
}

type valuesCtx struct {
	context.Context
	vals map[string]any
}

func value(c context.Context, key any) any {
	for {
		switch ctx := c.(type) {
		case *valuesCtx:
			kStr, _ := key.(string)
			if val, ok := ctx.vals[kStr]; ok {
				return val
			}
			c = ctx.Context
		default:
			return c.Value(key)
		}
	}
}

type stringer interface {
	String() string
}

func contextName(c context.Context) string {
	if s, ok := c.(stringer); ok {
		return s.String()
	}
	return reflect.TypeOf(c).String()
}

func (c *valuesCtx) String() string {
	vals, _ := json.Marshal(c.vals)

	return contextName(c.Context) + ".WithValue(type " +
		string(vals) + ")"
}

func (c *valuesCtx) Value(key any) any {
	kStr, _ := key.(string)
	if val, ok := c.vals[kStr]; ok {
		return val
	}
	return value(c.Context, key)
}
