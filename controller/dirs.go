package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-edge-bios/model"
	"github.com/gin-gonic/gin"
	"os"
)

type DirExistence struct {
	Path   string `json:"path"`
	Exists bool   `json:"exists"`
}

func (inst *Controller) DirExists(c *gin.Context) {
	path := c.Query("path")
	exists := fileutils.DirExists(path)
	dirExistence := DirExistence{Path: path, Exists: exists}
	responseHandler(dirExistence, nil, c)
}

func (inst *Controller) CreateDir(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		responseHandler(nil, errors.New("path can not be empty"), c)
		return
	}
	err := os.MkdirAll(path, os.FileMode(inst.FileMode))
	responseHandler(model.Message{Message: fmt.Sprintf("created directory: %s", path)}, err, c)
}
