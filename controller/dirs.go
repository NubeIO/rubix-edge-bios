package controller

import (
	"errors"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/files"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) DirExists(c *gin.Context) {
	path := c.Query("path")
	err := fileutils.DirExistsErr(path)
	var found bool
	if err == nil {
		found = true
	}
	reposeHandler(found, nil, c)
}

func (inst *Controller) CreateDir(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		reposeHandler(nil, errors.New("path can not be empty"), c)
		return
	}
	err := files.MakeDirectoryIfNotExists(path)
	reposeHandler(Message{Message: "directory creation is successfully executed"}, err, c)
}

func (inst *Controller) CopyDir(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	if from == "" || to == "" {
		reposeHandler(nil, errors.New("from and to directories name can not be empty"), c)
		return
	}
	exists := fileutils.DirExists(from)
	if !exists {
		reposeHandler(nil, errors.New("from dir not found"), c)
		return
	}
	err := fileutils.Copy(from, to)
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
	if !fileutils.DirExists(path) {
		reposeHandler(nil, errors.New("directory does not exist"), c)
		return
	}
	if recursively {
		err := fileutils.RmRF(path)
		if err != nil {
			reposeHandler(nil, err, c)
			return
		}
	} else {
		err := fileutils.Rm(path)
		if err != nil {
			reposeHandler(nil, err, c)
			return
		}
	}
	reposeHandler(Message{Message: "deletion of directory is successfully executed"}, nil, c)
	return
}
