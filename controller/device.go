package controller

import (
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetDeviceInfo(c *gin.Context) {
	deviceInfo, err := inst.RubixRegistry.GetDeviceInfo()
	responseHandler(deviceInfo, err, c)
}

func (inst *Controller) UpdateDeviceInfo(c *gin.Context) {
	var deviceInfo *rubixregistry.DeviceInfo
	err := c.ShouldBindJSON(&deviceInfo)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	response, err := inst.RubixRegistry.UpdateDeviceInfo(*deviceInfo)
	responseHandler(response, err, c)
}
