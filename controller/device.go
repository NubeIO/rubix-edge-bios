package controller

import (
	"github.com/NubeIO/rubix-edge/pkg/model"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetDeviceInfo(c *gin.Context) {
	data, err := inst.DB.GetDeviceInfo()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) UpdateDeviceInfo(c *gin.Context) {
	var m *model.DeviceInfo
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.DB.UpdateDeviceInfo(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) DropDeviceInfo(c *gin.Context) {
	host, err := inst.DB.DropDeviceInfo()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}
