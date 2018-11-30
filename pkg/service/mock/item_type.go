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

package mock

import (
	"context"
	"github.com/satori/go.uuid"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
)

func NewItemTypeService() *ItemTypeService {
	return &ItemTypeService{
		itemTypes: map[string]core.ItemType{
			"910da154-ba93-4a2e-bda4-83d45b212cf0": {Id: "910da154-ba93-4a2e-bda4-83d45b212cf0", Name: "Item Type 1", Description: "Item Type #1 in mock"},
			"de1162c8-4382-4822-ba4d-05a9a33b10a2": {Id: "de1162c8-4382-4822-ba4d-05a9a33b10a2", Name: "Item Type 2", Description: "Item Type #2 in mock"},
		},
	}
}

type ItemTypeService struct {
	itemTypes map[string]core.ItemType
}

func (s *ItemTypeService) List(ctx context.Context) (result []core.ItemType, err error) {
	result = make([]core.ItemType, 0, len(s.itemTypes))
	for _, itemType := range s.itemTypes {
		result = append(result, itemType)
	}
	return
}

func (s *ItemTypeService) Get(ctx context.Context, id string) (result *core.ItemType, err error) {
	if itemType, exists := s.itemTypes[id]; exists {
		return &itemType, nil
	}
	return nil, nil
}

func (s *ItemTypeService) Create(ctx context.Context, input *core.ItemType) (result *core.ItemType, err error) {
	result = input
	result.Id = uuid.NewV4().String()
	s.itemTypes[result.Id] = *result
	return
}

func (s *ItemTypeService) Update(ctx context.Context, id string, input *core.ItemType) (result *core.ItemType, err error) {
	if _, exists := s.itemTypes[id]; exists {
		input.Id = id
		s.itemTypes[id] = *input
		return input, nil
	}
	return nil, nil
}

func (s *ItemTypeService) Delete(ctx context.Context, id string) (err error) {
	if _, exists := s.itemTypes[id]; exists {
		delete(s.itemTypes, id)
		return nil
	}
	return core.ErrRecordNotExists
}
