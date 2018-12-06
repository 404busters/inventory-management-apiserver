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

	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
)

func BindUserServiceToContext(ctx context.Context, service core.UserService) context.Context {
	return withValue(ctx, UserServiceKey, service)
}

func GetUserServiceFromContext(ctx context.Context) core.UserService {
	val := ctx.Value(UserServiceKey)
	if service, ok := val.(core.UserService); ok {
		return service
	}
	return nil
}
