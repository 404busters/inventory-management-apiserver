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

package restful

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/service/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newLocationHandler() *locationHandler {
	return &locationHandler{
		Service: mock.NewMockLocationService(),
	}
}

func TestLocationHandler_list(t *testing.T) {
	handler := newLocationHandler()
	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	handler.list(c)
	if resp.Code != http.StatusOK {
		t.Errorf("incorrect status code, got %d", resp.Code)
	}
}
