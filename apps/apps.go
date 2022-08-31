package apps

import (
	"errors"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-edge-bios/constants"
)

type EdgeApps struct {
	App  *installer.App
	Perm int `json:"file_perm"`
	Ctl  *systemctl.Ctl
}

func New(apps *EdgeApps) (*EdgeApps, error) {
	if apps == nil {
		return nil, errors.New("store can not be empty")
	}
	if apps.App == nil {
		return nil, errors.New("app can not be empty")
	}
	if apps.App.FilePerm == 0 {
		apps.App.FilePerm = constants.Permission
	}
	if apps.App.DataDir == "" {
		apps.App.DataDir = "/data"
	}
	apps.App = installer.New(apps.App)
	apps.Ctl = systemctl.New(false, 30)
	return apps, nil
}
