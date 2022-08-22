package controller

import (
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) FlowToken(c *gin.Context) {
	sys := system.New(&system.System{})
	data, err := sys.GetFlowToken()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
