package apps

import (
	"errors"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
)

const nonRoot = 0700
const root = 0777

//var FilePerm = root

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
		apps.App.FilePerm = root
	}
	if apps.App.DataDir == "" {
		apps.App.DataDir = "/data"
	}
	apps.App = installer.New(apps.App)
	apps.Ctl = systemctl.New(&systemctl.Ctl{
		UserMode: false,
		Timeout:  30,
	})
	return apps, nil
}
