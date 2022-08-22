package controller

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/files"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) CreateDir(c *gin.Context) {
	path := c.Query("path")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	if path == "" {
		reposeHandler(nil, errors.New("path can not be empty"), c)
		return
	}
	err = files.MakeDirectoryIfNotExists(path)
	reposeHandler(Message{Message: "directory creation is successfully executed"}, err, c)
}

func (inst *Controller) CopyDir(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	if from == "" || to == "" {
		reposeHandler(nil, errors.New("from and to directories name can not be empty"), c)
		return
	}
	exists := fileUtils.DirExists(from)
	if !exists {
		reposeHandler(nil, errors.New("from dir not found"), c)
		return
	}
	err = fileUtils.Copy(from, to)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: "copying directory is successfully executed"}, err, c)
}

func (inst *Controller) DeleteDir(c *gin.Context) {
	path := c.Query("path")
	recursively := c.Query("recursively") == "true"
	if path == "" {
		reposeHandler(nil, errors.New("path can not be empty"), c)
		return
	}
	if !fileUtils.DirExists(path) {
		reposeHandler(nil, err, c)
		return
	}
	if recursively {
		err := fileUtils.RmRF(path)
		if err != nil {
			reposeHandler(nil, err, c)
			return
		}
	} else {
		err := fileUtils.Rm(path)
		if err != nil {
			reposeHandler(nil, err, c)
			return
		}
	}
	reposeHandler(Message{Message: "deletion of directory is successfully executed"}, nil, c)
	return
}
