package controller

import (
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) RunScanner(c *gin.Context) {
	var m *system.Scanner
	err = c.ShouldBindJSON(&m)
	data, err := inst.System.RunScanner(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
