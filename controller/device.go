package controller

import (
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetDeviceInfo(c *gin.Context) {
	deviceInfo, err := inst.RubixRegistry.GetDeviceInfo()
	reposeHandler(deviceInfo, err, c)
}

func (inst *Controller) UpdateDeviceInfo(c *gin.Context) {
	var deviceInfo *rubixregistry.DeviceInfo
	err := c.ShouldBindJSON(&deviceInfo)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	response, err := inst.RubixRegistry.UpdateDeviceInfo(*deviceInfo)
	reposeHandler(response, err, c)
}
