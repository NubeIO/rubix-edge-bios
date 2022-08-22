package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/gin-gonic/gin"
)

// CtlAction start, stop, enable, disable a service
func (inst *Controller) CtlAction(c *gin.Context) {
	var m *installer.CtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Rubix.App.CtlAction(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// CtlStatus check isRunning, isInstalled, isEnabled, isActive, isFailed for a service
func (inst *Controller) CtlStatus(c *gin.Context) {
	var m *installer.CtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Rubix.App.CtlStatus(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// ServiceMassAction start, stop, enable, disable a service
func (inst *Controller) ServiceMassAction(c *gin.Context) {
	var m *installer.CtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Rubix.App.ServiceMassAction(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// ServiceMassStatus on mass check isRunning, isInstalled, isEnabled, isActive, isFailed for a service
func (inst *Controller) ServiceMassStatus(c *gin.Context) {
	var m *installer.CtlBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Rubix.App.ServiceMassStatus(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
