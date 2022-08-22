package model

import (
	"github.com/NubeIO/lib-networking/networking"
	"github.com/NubeIO/lib-schema/schema"
)

type IpSchema struct {
	Interface schema.Interface `json:"interface"`
	Netmask   schema.Netmask   `json:"netmask"`
	Gateway   schema.Gateway   `json:"gateway"`
}

var nets = networking.New()

func GetIpSchema() *IpSchema {
	m := &IpSchema{}
	schema.Set(m)
	names, err := nets.GetInterfacesNames()
	if err != nil {
		return m
	}
	var out []string
	for _, name := range names.Names {
		if name != "lo" {
			out = append(out, name)
		}
	}
	m.Interface.Options = out
	return m
}
