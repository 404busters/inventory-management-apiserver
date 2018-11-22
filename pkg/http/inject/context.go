package inject

import (
	"context"
)

type injectContext struct {
	context.Context
	injects map[contextKey]interface{}
}

func (c *injectContext) Value(key interface{}) interface{} {
	if ctxKey, ok := key.(contextKey); ok {
		return c.injects[ctxKey]
	}
	return c.Context.Value(key)
}

func withValue(base context.Context, key contextKey, value interface{}) *injectContext {
	if injectCtx, ok := base.(*injectContext); ok {
		injectCtx.injects[key] = value
		return injectCtx
	}
	return &injectContext{
		Context: base,
		injects: map[contextKey]interface{}{
			key: value,
		},
	}
}
