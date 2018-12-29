package core

import "context"

type Preset struct {
	Id          string `json:"id,omitempty"`
	DisplayName string `json:"display_name"`
	ItemTypes   struct {
		Id    string `json:"id"`
		Count int    `json:"count"`
	} `graphql:"-"`
}

type PresetService interface {
	Create(ctx context.Context, preset *Preset) (result *Preset, err error)
	Get(ctx context.Context, id string) (preset *Preset, err error)
	Update(ctx context.Context, id string, preset *Preset) (result *Preset, err error)
	Delete(ctx context.Context, id string) (err error)
}
