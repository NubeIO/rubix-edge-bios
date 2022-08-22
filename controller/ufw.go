package controller

import (
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) UWFActive(c *gin.Context) {
	data, err := inst.System.UWFActive()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) UWFEnable(c *gin.Context) {
	data, err := inst.System.UWFEnable()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) UWFDisable(c *gin.Context) {
	data, err := inst.System.UWFDisable()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) UWFStatus(c *gin.Context) {
	data, err := inst.System.UWFStatus()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) UWFStatusList(c *gin.Context) {
	data, err := inst.System.UWFStatusList()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) UWFOpenPort(c *gin.Context) {
	var m system.UFWBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.System.UWFOpenPort(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) UWFClosePort(c *gin.Context) {
	var m system.UFWBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.System.UWFClosePort(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
