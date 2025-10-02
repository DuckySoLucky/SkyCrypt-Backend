package models

import skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"

type ItemTexture struct {
	Parent    string            `json:"parent"`
	Textures  map[string]string `json:"textures"`
	Overrides []Override        `json:"overrides"`
}

type Override struct {
	Predicate map[string]interface{} `json:"predicate"`
	Texture   string                 `json:"model"`
}

type TextureItem struct {
	Count   *int                                     `nbt:"Count" json:"Count,omitempty"`
	Damage  *int                                     `nbt:"Damage" json:"Damage,omitempty"`
	ID      *int                                     `nbt:"id" json:"id,omitempty"`
	Tag     skycrypttypes.TextureItemExtraAttributes `nbt:"tag" json:"tag,omitempty"`
	RawId   string                                   `nbt:"raw_id" json:"raw_id,omitempty"`
	Texture string                                   `nbt:"texture" json:"texture,omitempty"`
}

type VanillaTexture struct {
	VanillaId string `json:"vanillaId"`
	Damage    int    `json:"damage"`
}

type McMeta struct {
	Animation McMetaAnimation `json:"animation"`
}

type McMetaAnimation struct {
	Frametime int `json:"frametime"`
}
