package controller

import (
	"github.com/NubeIO/rubix-edge-bios/interfaces"
	"github.com/NubeIO/rubix-edge-bios/release"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) Ping(c *gin.Context) {
	responseHandler(interfaces.Ping{Version: release.GetVersion()}, nil, c)
}
