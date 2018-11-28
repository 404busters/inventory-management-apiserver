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

package core

import "context"

type Location struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type LocationService interface {
	List(ctx context.Context) ([]Location, error)
	Get(ctx context.Context, id string) (*Location, error)
	Create(ctx context.Context, input *Location) (*Location, error)
	Update(ctx context.Context, id string, input *Location) (*Location, error)
	Delete(ctx context.Context, id string) error
}
