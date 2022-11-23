package controller

import (
	"github.com/NubeIO/rubix-edge-bios/config"
	"github.com/NubeIO/rubix-edge-bios/model"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetArch(c *gin.Context) {
	deviceType := model.Arch{Arch: config.Config.GetDeviceType()}
	responseHandler(deviceType, nil, c)
}
