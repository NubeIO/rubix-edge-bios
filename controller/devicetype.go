package controller

import (
	"github.com/NubeIO/rubix-edge-bios/config"
	"github.com/NubeIO/rubix-edge-bios/model"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetDeviceType(c *gin.Context) {
	deviceType := model.DeviceType{DeviceType: config.Config.GetDeviceType()}
	responseHandler(deviceType, nil, c)
}
