package model

import "github.com/NubeIO/lib-schema/schema"

type UUID struct {
	Type     string `json:"type" default:"string"`
	ReadOnly string `json:"default" default:"true"`
}

type Name struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
	Min      int    `json:"min" default:"2"`
	Max      int    `json:"max" default:"50"`
	Default  string `json:"default" default:""`
}

type Description struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
	Min      int    `json:"min" default:"0"`
	Max      int    `json:"max" default:"80"`
}

type LocationUUID struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"true"`
	Binding  string `json:"binding" default:"locations/uuid"`
}

type NetworkSchema struct {
	UUID         UUID         `json:"uuid"`
	Name         Name         `json:"name"`
	Description  Description  `json:"description"`
	LocationUUID LocationUUID `json:"location_uuid"`
}

func GetNetworkSchema() *NetworkSchema {
	m := &NetworkSchema{}
	schema.Set(m)
	return m
}
