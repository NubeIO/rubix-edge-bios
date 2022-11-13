package controller

import (
	"github.com/NubeIO/rubix-edge-bios/model"
	"github.com/gin-gonic/gin"
	"syscall"
)

func (inst *Controller) SyscallUnlink(c *gin.Context) {
	path := c.Query("path")
	err := syscall.Unlink(path)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(model.Message{Message: "unlinked successfully"}, err, c)
}

func (inst *Controller) SyscallLink(c *gin.Context) {
	path := c.Query("path")
	link := c.Query("link")
	err := syscall.Symlink(path, link)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(model.Message{Message: "linked successfully"}, err, c)
}
