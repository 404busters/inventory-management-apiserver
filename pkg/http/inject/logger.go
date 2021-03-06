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

	"github.com/sirupsen/logrus"
)

func BindLoggerToContext(ctx context.Context, logger logrus.FieldLogger) context.Context {
	return withValue(ctx, LoggerKey, logger)
}

func GetLoggerFromContext(ctx context.Context) logrus.FieldLogger {
	val := ctx.Value(LoggerKey)
	if logger, ok := val.(logrus.FieldLogger); ok {
		return logger
	}

	return nil
}
