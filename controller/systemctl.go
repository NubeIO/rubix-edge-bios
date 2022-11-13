package controller

import (
	"github.com/NubeIO/lib-systemctl-go/systemctl/properties"
	"github.com/NubeIO/rubix-edge-bios/model"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) SystemCtlEnable(c *gin.Context) {
	unit := c.Query("unit")
	err := inst.SystemCtl.Enable(unit)
	message := model.Message{Message: "enabled successfully"}
	responseHandler(message, err, c)
}

func (inst *Controller) SystemCtlDisable(c *gin.Context) {
	unit := c.Query("unit")
	err := inst.SystemCtl.Disable(unit)
	message := model.Message{Message: "disabled successfully"}
	responseHandler(message, err, c)
}

func (inst *Controller) SystemCtlShow(c *gin.Context) {
	unit := c.Query("unit")
	p := c.Query("property")
	property, err := inst.SystemCtl.Show(unit, properties.Property(p))
	property_ := model.SystemCtlProperty{Property: property}
	responseHandler(property_, err, c)
}

func (inst *Controller) SystemCtlStart(c *gin.Context) {
	unit := c.Query("unit")
	err := inst.SystemCtl.Start(unit)
	message := model.Message{Message: "started successfully"}
	responseHandler(message, err, c)
}

func (inst *Controller) SystemCtlStatus(c *gin.Context) {
	unit := c.Query("unit")
	status, err := inst.SystemCtl.Status(unit)
	message := model.SystemCtlStatus{Status: status}
	responseHandler(message, err, c)
}

func (inst *Controller) SystemCtlStop(c *gin.Context) {
	unit := c.Query("unit")
	err := inst.SystemCtl.Stop(unit)
	message := model.Message{Message: "stopped successfully"}
	responseHandler(message, err, c)
}

func (inst *Controller) SystemCtlResetFailed(c *gin.Context) {
	err := inst.SystemCtl.RestartFailed()
	message := model.Message{Message: "reset-failed command executed successfully"}
	responseHandler(message, err, c)
}

func (inst *Controller) SystemCtlDaemonReload(c *gin.Context) {
	err := inst.SystemCtl.DaemonReload()
	message := model.Message{Message: "daemon-reload command executed successfully"}
	responseHandler(message, err, c)
}

func (inst *Controller) SystemCtlRestart(c *gin.Context) {
	unit := c.Query("unit")
	err := inst.SystemCtl.Restart(unit)
	message := model.Message{Message: "restarted successfully"}
	responseHandler(message, err, c)
}

func (inst *Controller) SystemCtlMask(c *gin.Context) {
	unit := c.Query("unit")
	err := inst.SystemCtl.Mask(unit)
	message := model.Message{Message: "masked successfully"}
	responseHandler(message, err, c)
}

func (inst *Controller) SystemCtlUnmask(c *gin.Context) {
	unit := c.Query("unit")
	err := inst.SystemCtl.Unmask(unit)
	message := model.Message{Message: "unmasked successfully"}
	responseHandler(message, err, c)
}

func (inst *Controller) SystemCtlState(c *gin.Context) {
	unit := c.Query("unit")
	state, err := inst.SystemCtl.State(unit)
	responseHandler(state, err, c)
}

func (inst *Controller) SystemCtlIsEnabled(c *gin.Context) {
	unit := c.Query("unit")
	state, err := inst.SystemCtl.IsEnabled(unit)
	state_ := model.SystemCtlState{State: state}
	responseHandler(state_, err, c)
}

func (inst *Controller) SystemCtlIsActive(c *gin.Context) {
	unit := c.Query("unit")
	state, status, err := inst.SystemCtl.IsActive(unit)
	status_ := model.SystemCtlStateStatus{State: state, Status: status}
	responseHandler(status_, err, c)
}

func (inst *Controller) SystemCtlIsRunning(c *gin.Context) {
	unit := c.Query("unit")
	state, status, err := inst.SystemCtl.IsRunning(unit)
	status_ := model.SystemCtlStateStatus{State: state, Status: status}
	responseHandler(status_, err, c)
}

func (inst *Controller) SystemCtlIsFailed(c *gin.Context) {
	unit := c.Query("unit")
	state, err := inst.SystemCtl.IsFailed(unit)
	state_ := model.SystemCtlState{State: state}
	responseHandler(state_, err, c)
}

func (inst *Controller) SystemCtlIsInstalled(c *gin.Context) {
	unit := c.Query("unit")
	state, err := inst.SystemCtl.IsInstalled(unit)
	state_ := model.SystemCtlState{State: state}
	responseHandler(state_, err, c)
}
