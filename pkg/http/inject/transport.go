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

func BindTransportServiceToContext(ctx context.Context, service core.TransportService) context.Context {
	return withValue(ctx, TransportServiceKey, service)
}

func GetTransportServiceFromContext(ctx context.Context) core.TransportService {
	val := ctx.Value(TransportServiceKey)
	if service, ok := val.(core.TransportService); ok {
		return service
	}
	return nil
}
