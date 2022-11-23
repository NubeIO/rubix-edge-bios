package controller

import (
	"github.com/NubeIO/rubix-edge-bios/config"
	"github.com/NubeIO/rubix-edge-bios/model"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetArch(c *gin.Context) {
	arch := model.Arch{Arch: config.Config.GetArch()}
	responseHandler(arch, nil, c)
}
