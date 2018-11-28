/*
	Copyright 2018 Carmen Chan & Tony Yip

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

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
